-- Create the actual table
CREATE TABLE public.poi
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    category JSONB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

-- Add location as a 2D POINT geometry type (ESRI compatible - SRID 4326)
SELECT AddGeometryColumn('public','poi','location','4326','POINT',2);

-- Add a new spatial index
CREATE INDEX location_idx ON public.poi USING GIST(location);

-- Update DB statistics after index creation
ANALYZE public.poi;
