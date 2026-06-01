DROP INDEX IF EXISTS affiliate_system.idx_offer_partner_id;
DROP INDEX IF EXISTS affiliate_system.idx_offer_category_id;
DROP INDEX IF EXISTS affiliate_system.idx_offer_expire_at;

DROP TABLE IF EXISTS affiliate_system.offer;
DROP TABLE IF EXISTS affiliate_system.partner;
DROP TABLE IF EXISTS affiliate_system.category;

DROP SCHEMA IF EXISTS affiliate_system;