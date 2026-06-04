DROP INDEX IF EXISTS affiliate_system.idx_offer_partner_id;
DROP INDEX IF EXISTS affiliate_system.idx_offer_category_id;
DROP INDEX IF EXISTS affiliate_system.idx_offer_city_id;
DROP INDEX IF EXISTS affiliate_system.idx_offer_expire_at;

DROP TABLE IF EXISTS affiliate_system.offers;
DROP TABLE IF EXISTS affiliate_system.cities;
DROP TABLE IF EXISTS affiliate_system.categories;
DROP TABLE IF EXISTS affiliate_system.partners;
DROP TABLE IF EXISTS affiliate_system.users;

DROP SCHEMA IF EXISTS affiliate_system CASCADE;