package buyer

const (
	GET_ALL_BUYERS  = "SELECT id, card_number_id, first_name, last_name FROM buyer"
	GET_BUYER_BY_ID = "SELECT id, card_number_id, first_name, last_name FROM buyer WHERE id=?"
	CREATE_BUYER    = "INSERT INTO buyer (card_number_id, first_name,last_name) VALUES(?,?,?,?)"
	UPDATE_BUYER    = "UPDATE buyer SET card_number_id=?, first_name=?, last_name=? WHERE id=?"
	DELETE_BUYER    = "DELETE FROM buyer WHERE id=?"
)
