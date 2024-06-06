CREATE OR REPLACE FUNCTION filter_records(
    search_name VARCHAR(100),
    search_descr VARCHAR(200),
    search_quant INT,
    search_price NUMERIC(6,2),
    search_topic VARCHAR(100),
    search_created DATE,
    search_published TIMESTAMP,
    search_page INT,
    search_limit INT
)
RETURNS TABLE (
    name VARCHAR(100),
    description VARCHAR(200),
    record_image TEXT,
    record_price NUMERIC(6,2),
    record_topic VARCHAR(100),
    record_created DATE,
    record_published TIMESTAMP,
    record_quant INT
)
AS $$
DECLARE
    max_records INT := search_limit;  
    offset_val INT := (search_page - 1) * search_limit;  
BEGIN
    RETURN QUERY
    SELECT record_name, descr, img, price, topic, created, published, quant
    FROM record
    WHERE
        (search_name IS NULL OR record.record_name ILIKE '%' || search_name || '%')
        AND (search_descr IS NULL OR to_tsvector('spanish', record.descr) @@ plainto_tsquery('spanish', search_descr))
        AND (search_price IS NULL OR record.price::TEXT ILIKE '%' || search_price || '%')
        AND (search_topic IS NULL OR record.topic =search_topic)
        AND (search_created IS NULL OR record.created = search_created)
        AND (search_published IS NULL OR 
             EXTRACT(YEAR FROM record.published) = EXTRACT(YEAR FROM search_published) OR
             EXTRACT(MONTH FROM record.published) = EXTRACT(MONTH FROM search_published) OR
             EXTRACT(DAY FROM record.published) = EXTRACT(DAY FROM search_published))
        AND (search_quant IS NULL OR record.quant::TEXT ILIKE '%' || search_quant || '%')
    LIMIT max_records
    OFFSET offset_val;  
END;
$$ LANGUAGE plpgsql;
