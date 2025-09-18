-- Migration 002: Create Causal Relations Table
-- This migration creates the causal_relations table for storing Aristotelian causes:
-- - Material Cause (what something is made of)
-- - Formal Cause (the essence or form)
-- - Efficient Cause (what brings it about)
-- - Final Cause (the purpose or end)

-- Create causal_relations table (moved from 001 to avoid dependency issues)
CREATE TABLE causal_relations (
    id TEXT PRIMARY KEY,
    cause_type TEXT NOT NULL,
    from_entity TEXT NOT NULL,
    to_entity TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
    -- Note: Foreign keys removed to allow references to any entity type
    -- This allows causal relations between substances, kinds, attributes, etc.
);

-- Create indexes for causal relations
CREATE INDEX idx_causal_relations_from_entity ON causal_relations(from_entity);
CREATE INDEX idx_causal_relations_to_entity ON causal_relations(to_entity);
CREATE INDEX idx_causal_relations_cause_type ON causal_relations(cause_type);
