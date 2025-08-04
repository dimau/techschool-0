-- Rollback Migration: Drop orders tables
-- Description: Drops all tables created for order storage

-- Drop triggers first
DROP TRIGGER IF EXISTS update_items_updated_at ON items;
DROP TRIGGER IF EXISTS update_payment_updated_at ON payment;
DROP TRIGGER IF EXISTS update_delivery_updated_at ON delivery;
DROP TRIGGER IF EXISTS update_orders_updated_at ON orders;

-- Drop trigger function
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop tables in reverse order (due to foreign key constraints)
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS delivery;
DROP TABLE IF EXISTS orders; 