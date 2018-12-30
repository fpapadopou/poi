-- Create the actual table
CREATE TABLE public.poi
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    longitude DOUBLE PRECISION,
    latitude DOUBLE PRECISION,
    category JSONB,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- Add location as a 2D POINT geometry type (ESRI compatible - SRID 4326)
-- Column is added with AddGeometryColumn function in order to allow usage of OpenGIS support functions
SELECT AddGeometryColumn('public','poi','location','4326','POINT',2);

-- Add a new spatial index
CREATE INDEX location_idx ON public.poi USING GIST(location);

-- Update DB statistics after index creation
ANALYZE public.poi;
