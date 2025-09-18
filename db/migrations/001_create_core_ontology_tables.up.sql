-- Migration 001: Create Core Ontology Tables
-- This migration creates the foundational tables for E.J. Lowe's Four-Category Ontology:
-- 1. Kinds (natural classifications)
-- 2. Attributes (general properties) 
-- 3. Substances (independent entities)
-- 4. Modes (particular instantiations)
-- Plus supporting tables for potentialities, actualities, and relationships

-- Create kinds table
CREATE TABLE kinds (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL
);

-- Create attributes table
CREATE TABLE attributes (
    id TEXT PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    description TEXT,
    data_type TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Create substances table
CREATE TABLE substances (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    kind TEXT NOT NULL,
    essence TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Create modes table
CREATE TABLE modes (
    id TEXT PRIMARY KEY,
    value TEXT NOT NULL,
    substance_id TEXT NOT NULL,
    attribute_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (substance_id) REFERENCES substances(id) ON DELETE CASCADE,
    FOREIGN KEY (attribute_id) REFERENCES attributes(id) ON DELETE CASCADE
);

-- Create substance_attributes junction table
CREATE TABLE substance_attributes (
    substance_id TEXT NOT NULL,
    attribute_id TEXT NOT NULL,
    PRIMARY KEY (substance_id, attribute_id),
    FOREIGN KEY (substance_id) REFERENCES substances(id) ON DELETE CASCADE,
    FOREIGN KEY (attribute_id) REFERENCES attributes(id) ON DELETE CASCADE
);

-- Create potentialities table
CREATE TABLE potentialities (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    conditions TEXT, -- JSON string of required conditions
    substance_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    FOREIGN KEY (substance_id) REFERENCES substances(id) ON DELETE CASCADE
);

-- Create actualities table
CREATE TABLE actualities (
    id TEXT PRIMARY KEY,
    description TEXT NOT NULL,
    actualized_at TIMESTAMP NOT NULL,
    substance_id TEXT NOT NULL,
    potentiality_id TEXT NOT NULL,
    FOREIGN KEY (substance_id) REFERENCES substances(id) ON DELETE CASCADE,
    FOREIGN KEY (potentiality_id) REFERENCES potentialities(id) ON DELETE CASCADE
);

-- Create indexes
CREATE INDEX idx_kinds_name ON kinds(name);
CREATE INDEX idx_attributes_name ON attributes(name);
CREATE INDEX idx_substances_name ON substances(name);
CREATE INDEX idx_substances_kind ON substances(kind);
CREATE INDEX idx_substances_created_at ON substances(created_at);
CREATE INDEX idx_modes_substance_id ON modes(substance_id);
CREATE INDEX idx_modes_attribute_id ON modes(attribute_id);
CREATE INDEX idx_potentialities_substance_id ON potentialities(substance_id);
CREATE INDEX idx_actualities_substance_id ON actualities(substance_id);
CREATE INDEX idx_actualities_potentiality_id ON actualities(potentiality_id);

-- Insert sample data for testing
INSERT INTO kinds (id, name, description, created_at) VALUES 
('kind-1', 'Oak', 'A type of tree with potentiality for growth and leaf production', NOW()),
('kind-2', 'Human', 'Rational animal with potentiality for virtue and wisdom', NOW()),
('kind-3', 'Pine', 'Evergreen coniferous tree with potentiality for cone production', NOW());

INSERT INTO attributes (id, name, description, data_type, created_at) VALUES 
('attr-1', 'color', 'Visual property of substances', 'string', NOW()),
('attr-2', 'virtue', 'Moral excellence and character', 'string', NOW()),
('attr-3', 'height', 'Vertical measurement of substances', 'number', NOW()),
('attr-4', 'health', 'Physical condition of substances', 'string', NOW());

-- Note: Substances, modes, potentialities, and actualities will be created via API calls
