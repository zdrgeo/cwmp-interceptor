# CWMP interceptor

## Overview

This repository explores an approach to collecting telemetry from CPE devices. It contains anoter "frontend" for the CWMP and USP [Bulk Data Collector](https://github.com/zdrgeo/bulk-data-collector).
Instead of accepting bulk data reports in CSV or JSON format, this "frontend" intercepts the communication between the CPEs and the ACS. Eavesdrop the CWMP protocol Inform messages. Synthesize telemetry events from the collected device parameters and other device events and sends them to a dedicated analytics or telemetry platform using the available "backend" implementations from [Bulk Data Collector](https://github.com/zdrgeo/bulk-data-collector).

## Context

CWMP Intercepting Collector

```mermaid
graph LR
    I((Interceptor))
    CPE1[CPE 1] <-.->|commands| I
    CPEn[CPE N] <-.->|commands| I
    I <-.->|commands| ACS[ACS]
    subgraph IC[Intercepting Collector]
        I -->|data point events| C((Collector))
    end
```

CWMP Eavesdropping Collector

```mermaid
graph LR
    RP[Reverse Proxy with Request Mirroring]
    CPE1[CPE 1] <-.->|commands| RP
    CPEn[CPE N] <-.->|commands| RP
    RP <-.->|commands| ACS[ACS]
    RP -->|commands| E((Eavesdropper))
    subgraph EC[Eavesdropping Collector]
        E -->|data point events| C((Collector))
    end
```
