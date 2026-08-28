CREATE INDEX IF NOT EXISTS idx_detalle_preliquidacion_contrato_preliquidacion
    ON liquidador.detalle_preliquidacion (contrato_preliquidacion_id);
