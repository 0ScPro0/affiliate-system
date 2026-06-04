ALTER TABLE affiliate_system.offer 
DROP CONSTRAINT IF EXISTS fk_offer_city;

ALTER TABLE affiliate_system.offer 
DROP COLUMN IF EXISTS city_id;

DROP TABLE IF EXISTS affiliate_system.city CASCADE;