-- Rename back: partners -> partner, categories -> category, offers -> offer, cities -> city

ALTER TABLE IF EXISTS affiliate_system.offers RENAME CONSTRAINT fk_offers_city TO fk_offer_city;

ALTER INDEX IF EXISTS affiliate_system.idx_offers_expire_at RENAME TO idx_offer_expire_at;
ALTER INDEX IF EXISTS affiliate_system.idx_offers_category_id RENAME TO idx_offer_category_id;
ALTER INDEX IF EXISTS affiliate_system.idx_offers_partner_id RENAME TO idx_offer_partner_id;

ALTER TABLE IF EXISTS affiliate_system.cities RENAME TO city;
ALTER TABLE IF EXISTS affiliate_system.offers RENAME TO offer;
ALTER TABLE IF EXISTS affiliate_system.categories RENAME TO category;
ALTER TABLE IF EXISTS affiliate_system.partners RENAME TO partner;