INSERT INTO categories (id, name, created_at, updated_at, deleted_at)
VALUES
    (uuid_generate_v4(), 'Electronics', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (uuid_generate_v4(), 'Clothing', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (uuid_generate_v4(), 'Health & Beauty', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (uuid_generate_v4(), 'Home & Kitchen', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
    (uuid_generate_v4(), 'Sports & Outdoors', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);

INSERT INTO products (id, name, description, category_id, price, stock, rating, created_at, updated_at, deleted_at)
VALUES
    (uuid_generate_v4(), 'Smartphone X1', 'Smartphone dengan layar 6.5 inci, RAM 4GB, dan penyimpanan 64GB', '0452e5f3-58a6-4a74-b9d7-851fa3d6d9d8', 4999000.00, 50, 4.5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL), -- Electronics
    (uuid_generate_v4(), 'T-Shirt Basic', 'T-Shirt berbahan katun dengan berbagai pilihan warna', '2f7a42fa-457f-426d-a4b4-236d83a1e514', 150000.00, 100, 4.2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL), -- Clothing
    (uuid_generate_v4(), 'Shampoo Herbal', 'Shampoo dengan bahan alami untuk perawatan rambut', '556f6a72-bd26-408d-9c5b-0ed832341d03', 75000.00, 200, 4.3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL), -- Health & Beauty
    (uuid_generate_v4(), 'Electric Kettle', 'Ketel listrik dengan kapasitas 1.7 liter dan fitur pemanas cepat', '29cbe979-554f-4a41-8ca7-f5e63a6b5923', 350000.00, 30, 4.7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL), -- Home & Kitchen
    (uuid_generate_v4(), 'Camping Tent', 'Tenda camping dengan 4 tempat tidur dan pelindung cuaca', 'a5fff2b4-75f4-4832-bf97-0cef8f4816b7', 1400000.00, 20, 4.6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL); -- Sports & Outdoors
