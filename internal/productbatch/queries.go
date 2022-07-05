package productbatch


const (
	sqlCreatePB = `insert into mercadofresco.products_batches (batch_number,current_quantity
			,current_temperature,due_date,initial_quantity,manufacturing_date,manufacturing_hour,minimum_temperature,product_id,section_id)
			values(?,?,?,?,?,?,?,?,?,?)`
	sqlBatcheNumber = `SELECT batch_number FROM mercadofresco.products_batches WHERE batch_number=?`
	sqlrelatorioTodo = `SELECT s.id AS 'section_id', s.section_number AS 'section_number', count(*) AS 'products_count'
			FROM mercadofresco.products_batches as pbs
			INNER JOIN mercadofresco.sections AS s ON s.id = pbs.section_id 
			INNER JOIN mercadofresco.products AS pcts ON pbs.product_id = pcts.id
			GROUP BY s.id`
	sqlrelatorioSectioId = `SELECT s.id as 'section_id', s.section_number, count(*) as 'products_count'
			from mercadofresco.products_batches as pbs
			INNER JOIN mercadofresco.sections as s on s.id = pbs.section_id 
			INNER JOIN mercadofresco.products as pcts on pbs.product_id = pcts.id
			WHERE s.id=?`
)
 
 