# Tailscale ACL CLI — TODO Roadmap

## 1. MVP
- [ ] **Auth**
  - [ ] OAuth device flow
  - [ ] API key auth
  - [ ] Secure token storage (OS keyring / file 0600 fallback)
  - [ ] Least-privilege scopes + overscope warning
- [ ] **Core ACL ops**
  - [ ] `acl get` (fetch current ACL)
  - [ ] `acl set` (apply from file/stdin)
  - [ ] `acl validate` (schema check)
  - [ ] `acl lint` (basic rules)
  - [ ] `--dry-run` support everywhere
- [ ] **Safety**
  - [ ] Lockout guard (detect removing own access)
- [ ] **Ergonomics**
  - [ ] Helpful errors, structured logs
  - [ ] Shell completion

---

## 2. Professional-Grade
- [ ] **Plan / Apply workflow**
  - [ ] `acl plan` (show diff, exit 2 on changes)
  - [ ] `acl apply --approve`
  - [ ] Human + machine-readable diffs
- [ ] **Policy linting**
  - [ ] Org-configurable rules (YAML)
  - [ ] Severity levels (error/warn/info)
- [ ] **Schema validation**
  - [ ] Versioned ACL schema support
- [ ] **Drift detection**
  - [ ] `drift scan` (compare remote vs Git HEAD)
- [ ] **Resiliency**
  - [ ] Exponential backoff, retry, idempotent ops
- [ ] **Observability**
  - [ ] `--trace` flag (redacted request/response)
  - [ ] Opt-in anonymous metrics

---

## 3. Security & Guardrails
- [ ] **Change provenance**
  - [ ] Annotate changes (`--reason`, `GIT_COMMIT`)
- [ ] **Signed policy bundles**
  - [ ] SOPS/age or GPG signing
- [ ] **RBAC mapping**
  - [ ] Enforce roles in CLI
- [ ] **Break-glass**
  - [ ] Revert to last known good ACL

---

## 4. GitOps & CI
- [ ] **CI mode**
  - [ ] Non-interactive, deterministic, exit codes
- [ ] **Pre-commit / conftest**
  - [ ] Run validate + lint before commits
  - [ ] Support custom OPA/Rego rules
- [ ] **Examples**
  - [ ] GitHub Actions / GitLab CI / Jenkins pipelines

---

## 5. Tailscale-Specific Superpowers
- [ ] **Entity resolution**
  - [ ] Validate users, groups, tagged-devices exist
- [ ] **Impact analysis**
  - [ ] List affected nodes/users by policy change
- [ ] **Access query**
  - [ ] `access simulate --user X --dest Y:Z`
- [ ] **Environment overlays**
  - [ ] Policy fragments for `dev|staging|prod`

---

## 6. UX Polish
- [ ] **TUI mode (optional)**
  - [ ] Bubble Tea-based editor with diff + validation
- [ ] **Docs**
  - [ ] Quickstart, FAQ, risk notes
  - [ ] `--example` commands
- [ ] **Errors that teach**
  - [ ] Link to docs in error messages

---

## 7. Testing & Quality
- [ ] **Unit tests** (auth, parsing, diff, validation)
- [ ] **Integration tests** (fake Tailscale API, flakiness)
- [ ] **Load tests** (large ACLs, report timings)

---

## 8. Release & Distribution
- [ ] **goreleaser**
  - [ ] Multi-arch binaries
  - [ ] Homebrew, Scoop, AUR
  - [ ] Checksums + SBOM
- [ ] **Versioning**
  - [ ] SemVer + machine-readable `version`
- [ ] **Telemetry policy**
  - [ ] Default off, clear opt-in

---

## 9. Enterprise Features (Stretch)
- [ ] **Policy library**
  - [ ] Reusable snippets with parameters
- [ ] **Approval workflow**
  - [ ] `plan` posts PR comment with diff
  - [ ] `apply` gated by PR approval
- [ ] **Multi-tenant orgs**
  - [ ] Profiles for multiple orgs, isolated creds
