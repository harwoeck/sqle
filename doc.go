/*
Package sqle is a general purpose, transparent, non-magical helper package
for sql.DB that simplifies and reduces error checking for various operations.

In order to initialize a new `*sqle.Sqle` instance you need an working
`*sql.DB` connection. Pass it to the `New` function

	var db *sql.DB
	db := myDBInitAndOpen()

	var s *sqle.Sqle
	s = sqle.New(db)

If you have your `*sqle.Sqle` instance you can implement functions similar
to this

	var (
		s *sqle.Sqle
	)

	// InsertMessage inserts the passed string into the msg table and
	// returns the auto-incremented id (foreign-key in other tables).
	func InsertMessage(msg string) (id int64, err error) {
		return s.ExecID(queryInsertMsg, msg)
	}

	// SelectMessage receives the message associated to the passed
	// id.
	func SelectMessage(id int64) (msg string, exists bool, err error) {
		exists, err = s.SelectExists("SELECT message FROM msg WHERE id=?", []interface{}{id}, []interface{}{&msg})
		return
	}

	// ExistsMessage checks whether the id (and therefore) it's
	// associated value exist or not.
	func ExistsMessage(id int64) (exists bool, err error) {
		return s.Exists("SELECT id FROM msg WHERE id=?", id)
	}

	// SelectMessageExists checks whether the id (and therefore) it's
	// associated value exists or not and if so also returns the value.
	// If the id doesn't exists the msg pointer will be nil.
	func SelectMessageExists(id int64) (msg *string, err error) {
		exists, err := s.SelectExists("SELECT message FROM msg WHERE id=?", []interface{}{id}, []interface{}{msg})
		if err != nil || !exists {
			return nil, err
		}
		return msg, err
	}

More sophisticated usages of `SelectRange` could look like the following (these
examples are copied from the vikebot.com database code)

	// JoinedUsersCtx returns the `userID`s of all users which joined the round
	// specified through the roundID.
	func JoinedUsersCtx(roundID int, ctx *zap.Logger) (joined []int, success bool) {
		users := []int{}

		var userID int
		err := s.SelectRange("SELECT user_id FROM roundentry WHERE round_id=?",
			[]interface{}{roundID},
			[]interface{}{&userID},
			func() {
				users = append(users, userID)
			})
		if err != nil {
			ctx.Error("vbdb.JoinedUsersCtx",
				zap.Int("roundID", roundID),
				zap.Error(err))
			return nil, false
		}

		return users, true
	}

	// ActiveRoundsCtx loads all rounds which have not status `vbcore.RoundStatusFinished`.
	func ActiveRoundsCtx(ctx *zap.Logger) (rounds []vbcore.Round, success bool) {
		rounds = []vbcore.Round{}

		var id, joined, min, max, roundstatus int
		var name, wallpaper string
		var starttime mysql.NullTime
		err := s.SelectRange(`
			SELECT round.id,
				round.name,
				round.wallpaper,
				(SELECT COUNT(id) FROM roundentry WHERE roundentry.round_id = round.id) AS "joined",
				roundsize.min,
				roundsize.max,
				round.starttime,
				round.roundstatus_id
			FROM round, roundsize
			WHERE round.roundstatus_id IN (?, ?, ?)
			ORDER BY round.id ASC`,
			[]interface{}{vbcore.RoundStatusOpen, vbcore.RoundStatusClosed, vbcore.RoundStatusRunning},
			[]interface{}{&id, &name, &wallpaper, &joined, &min, &max, &starttime, &roundstatus},
			func() {
				r := vbcore.Round{
					ID:          id,
					Name:        name,
					Wallpaper:   wallpaper,
					Joined:      joined,
					Min:         min,
					Max:         max,
					Starttime:   starttime.Time,
					RoundStatus: roundstatus,
				}
				rounds = append(rounds, r)
			})
		if err != nil {
			ctx.Error("vbdb.ActiveRoundsCtx", zap.Error(err))
			return nil, false
		}

		return rounds, true
	}
*/
package sqle
