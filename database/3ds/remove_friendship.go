package database_3ds

import (
	"github.com/PretendoNetwork/friends-secure/database"
)

// RemoveFriendship removes a user's friend relationship
func RemoveFriendship(user1_pid uint32, user2_pid uint32) error {
	_, err := database.Postgres.Exec(`
		DELETE FROM "3ds".friendships WHERE user1_pid=$1 AND user2_pid=$2`, user1_pid, user2_pid)
	if err != nil {
		return err
	}

	_, err = database.Postgres.Exec(`
		UPDATE "3ds".friendships SET type=0 WHERE user1_pid=$1 AND user2_pid=$2`, user2_pid, user1_pid)
	if err != nil {
		return err
	}

	return nil
}
