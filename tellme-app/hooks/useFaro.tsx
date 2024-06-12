'use client';
import { useRef, useEffect } from 'react';
import { FetchInstrumentation } from '@grafana/faro-instrumentation-fetch';
import { XHRInstrumentation } from '@grafana/faro-instrumentation-xhr';
import { TracingInstrumentation } from '@grafana/faro-web-tracing';
import { JaegerPropagator } from '@opentelemetry/propagator-jaeger';
import { ZoneContextManager } from '@opentelemetry/context-zone';
import {
  initializeFaro as coreInit,
  getWebInstrumentations,
} from '@grafana/faro-react';

import type { Faro } from '@grafana/faro-react';

function initializeFaro(): Faro {
  const instrumentationOptions = {
    propagateTraceHeaderCorsUrls: [
      new RegExp(process.env.NEXT_PUBLIC_BACKEND_URL || ''),
    ],
  };
  const propagator = new JaegerPropagator();
  const contextManager = new ZoneContextManager();
  const faro = coreInit({
    url: process.env.NEXT_PUBLIC_FARO_COLLECT_ENDPOINT,
    apiKey: process.env.NEXT_PUBLIC_FARO_API_KEY,
    app: {
      name: 'learning-o11y-app',
      version: '0.1.0',
    },
    logArgsSerializer: (args) =>
      args.map((arg) => JSON.stringify(arg)).join(' '),
    instrumentations: [
      ...getWebInstrumentations({
        captureConsole: true,
        captureConsoleDisabledLevels: [],
      }),
      new FetchInstrumentation({}),
      new XHRInstrumentation({}),
      new TracingInstrumentation({
        instrumentationOptions,
        propagator,
        contextManager,
      }),
    ],
  });

  faro.api.pushLog(['Faro was initialized']);

  return faro;
}

export default function useFaro() {
  const faroRef = useRef<Faro | null>(null);
  useEffect(() => {
    if (!faroRef.current) {
      faroRef.current = initializeFaro();
    }
  }, []);

  return { faro: faroRef.current };
}
