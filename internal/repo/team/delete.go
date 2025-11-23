package team

import "context"

func (r *teamRepo) Delete(ctx context.Context, name string) error {
	// Сначала обнуляем team_name у пользователей
	updateUsersQuery := `UPDATE users SET team_name = '' WHERE team_name = $1`
	_, err := r.db.ExecContext(ctx, updateUsersQuery, name)
	if err != nil {
		return err
	}

	// Затем удаляем саму команду
	deleteTeamQuery := `DELETE FROM teams WHERE name = $1`
	_, err = r.db.ExecContext(ctx, deleteTeamQuery, name)
	return err
}
