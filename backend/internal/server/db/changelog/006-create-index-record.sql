CREATE INDEX idx_record_name ON record(record_name);
CREATE INDEX idx_record_descr ON record USING gin(to_tsvector('spanish', descr));
CREATE INDEX idx_record_price ON record(price);
CREATE INDEX idx_record_topic ON record(topic);
CREATE INDEX idx_record_created ON record(created);
CREATE INDEX idx_record_published ON record(published);
CREATE INDEX idx_record_quant ON record(quant);