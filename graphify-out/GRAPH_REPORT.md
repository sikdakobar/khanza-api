# Graph Report - sikdakobar__khanza-api  (2026-11-09)

> Cakupan: kode (AST) saja; dokumen belum diekstrak. Lihat `STATUS.md`.

## Corpus Check
- 19 files · ~2,759 words
- Verdict: corpus is large enough that graph structure adds value.

## Summary
- 121 nodes · 227 edges · 18 communities (14 shown, 4 thin omitted)
- Extraction: 82% EXTRACTED · 18% INFERRED · 0% AMBIGUOUS · INFERRED: 41 edges (avg confidence: 0.8)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `631af841`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- MongoDB
- coa.go
- pegawai.go
- obat.go
- harga.go
- config.go
- antrian.go
- Index
- Pasien
- poli.go
- Changelog
- main
- AGENTS.md
- README.md
- simpus

## God Nodes (most connected - your core abstractions)
1. `MongoDB()` - 42 edges
2. `Pasien` - 5 edges
3. `Index()` - 4 edges
4. `Store()` - 4 edges
5. `Update()` - 4 edges
6. `Destroy()` - 4 edges
7. `ChartAccount` - 4 edges
8. `IndexCOA()` - 4 edges
9. `ShowCOA()` - 4 edges
10. `StoreCOA()` - 4 edges

## Surprising Connections (you probably didn't know these)
- `Index()` --calls--> `MongoDB()`  [INFERRED]
  app/harga/harga.go → db/db.go
- `Store()` --calls--> `MongoDB()`  [INFERRED]
  app/harga/harga.go → db/db.go
- `Update()` --calls--> `MongoDB()`  [INFERRED]
  app/harga/harga.go → db/db.go
- `Destroy()` --calls--> `MongoDB()`  [INFERRED]
  app/harga/harga.go → db/db.go
- `IndexCOA()` --calls--> `MongoDB()`  [INFERRED]
  app/keuangan/coa.go → db/db.go

## Import Cycles
- None detected.

## Communities (18 total, 4 thin omitted)

### Community 0 - "MongoDB"
Cohesion: 0.36
Nodes (13): AlamatStore(), AlamatUpdate(), Destroy(), Request, ResponseWriter, Index(), RekamMedisStore(), RekamMedisUpdate() (+5 more)

### Community 1 - "coa.go"
Cohesion: 0.29
Nodes (13): DestroyCOA(), DestroySubAccountCOA(), DateTime, ObjectID, Request, ResponseWriter, IndexCOA(), ShowCOA() (+5 more)

### Community 2 - "pegawai.go"
Cohesion: 0.29
Nodes (13): Destroy(), DateTime, ObjectID, Request, ResponseWriter, Index(), Show(), Store() (+5 more)

### Community 3 - "obat.go"
Cohesion: 0.32
Nodes (12): BatchObatDestroy(), BatchObatStore(), BatchObatUpdate(), Destroy(), ObjectID, Request, ResponseWriter, Index() (+4 more)

### Community 4 - "harga.go"
Cohesion: 0.33
Nodes (9): Destroy(), DateTime, ObjectID, Request, ResponseWriter, Index(), Store(), Update() (+1 more)

### Community 5 - "config.go"
Cohesion: 0.42
Nodes (9): ConnectDB(), CreateUser(), GetUser(), Database, Request, ResponseWriter, IndexCollection(), IndexDB() (+1 more)

### Community 6 - "antrian.go"
Cohesion: 0.33
Nodes (8): Antrian, DateTime, ObjectID, Request, ResponseWriter, Index(), ListAntrian(), Store()

### Community 7 - "Index"
Cohesion: 0.32
Nodes (7): Apotek, DateTime, ObjectID, Request, ResponseWriter, Index(), Store()

### Community 8 - "Pasien"
Cohesion: 0.47
Nodes (6): DateTime, ObjectID, Alamat, Biometrik, Pasien, Rekam_Medis

### Community 9 - "poli.go"
Cohesion: 0.47
Nodes (5): Request, ResponseWriter, Index(), Store(), Poli

### Community 10 - "Changelog"
Cohesion: 0.50
Nodes (3): Added, Changelog, [v0.1.0] - 2026-08-22

## Knowledge Gaps
- **5 isolated node(s):** `Poli`, `simpus`, `Knowledge Graph (graphify)`, `Added`, `khanza-api`
  These have ≤1 connection - possible missing edges or undocumented components.
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `MongoDB()` connect `MongoDB` to `coa.go`, `pegawai.go`, `obat.go`, `harga.go`, `config.go`, `antrian.go`, `Index`?**
  _High betweenness centrality (0.590) - this node is a cross-community bridge._
- **Why does `Index()` connect `Index` to `MongoDB`?**
  _High betweenness centrality (0.090) - this node is a cross-community bridge._
- **Are the 40 inferred relationships involving `MongoDB()` (e.g. with `Destroy()` and `Index()`) actually correct?**
  _`MongoDB()` has 40 INFERRED edges - model-reasoned connections that need verification._
- **What connects `Poli`, `simpus`, `Knowledge Graph (graphify)` to the rest of the system?**
  _5 weakly-connected nodes found - possible documentation gaps or missing edges._