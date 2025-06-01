# CWMP interceptor

## Overview

Intercepts the communication between the CPEs and the ACS. Eavesdrop the CWMP protocol Inform messages.

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
