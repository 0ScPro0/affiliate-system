-- Rename tables: partner -> partners, category -> categories, offer -> offers, city -> cities
-- PostgreSQL automatically updates foreign key references when renaming tables

ALTER TABLE IF EXISTS affiliate_system.partner RENAME TO partners;
ALTER TABLE IF EXISTS affiliate_system.category RENAME TO categories;
ALTER TABLE IF EXISTS affiliate_system.offer RENAME TO offers;
ALTER TABLE IF EXISTS affiliate_system.city RENAME TO cities;

-- Rename indexes to match new table names
ALTER INDEX IF EXISTS affiliate_system.idx_offer_partner_id RENAME TO idx_offers_partner_id;
ALTER INDEX IF EXISTS affiliate_system.idx_offer_category_id RENAME TO idx_offers_category_id;
ALTER INDEX IF EXISTS affiliate_system.idx_offer_expire_at RENAME TO idx_offers_expire_at;

-- Rename foreign key constraint
ALTER TABLE IF EXISTS affiliate_system.offers RENAME CONSTRAINT fk_offer_city TO fk_offers_city;