DevSecOps Senior Interview Plan (60 ± 5 mins)

Role Focus: Senior DevSecOps Engineer (5+ years). Core emphasis on: Git fluency, CI/CD architecture (GitHub Actions / GitLab CI), secure software delivery (SSDLC), security scanners (Secrets, SAST, SCA, IAC), vulnerability triage, code review capability (Go/Python), Kubernetes operational security awareness, proactive mindset, and ability to reason under ambiguity.


---

High-Level Structure (Target: 60 mins + up to 5 mins buffer)

Segment
Duration
Format
Primary Objectives
1. Intro & Context
5 min
Conversational
Rapport, calibration, candidate framing of experience
2. Git & Repo Handling
8 min
Hands-on + Q&A
Branching, cloning, commits, hooks, remote strategy
3. Scanner Result Interpretation
12 min
Scenario discussion
Triage, false positive handling, mitigation strategy
4. Pipeline Extension + Debug
15 min
Practical design + verbal walkthrough
Add SCA + IAC scanner, fix broken CI, optimize
5. Code Review & Secure Fix
10 min
Hands-on
Spot vuln, propose fix, secure coding reasoning
6. Kubernetes & SSDLC Integration
5 min
Q&A
Secure deployment lifecycle, K8 misconfig reasoning
7. Incident & Zero-Day Response
5 min
Scenario-based
Log4Shell-style rapid response plan
8. Wrap-up / Candidate Questions
~5 min
Conversational
Assess curiosity, proactive mindset

Total (core): 60 min  
Buffer (overhead / transitions): +5 min


---

Improvements Over Previous Draft

Enhancements added versus initial draft:
- Dedicated code review block (secure refactor reasoning)
- Kubernetes + SSDLC consolidation for efficiency
- Incident response / zero‑day handling
- Structured scoring rubric with explicit performance tiers
- Clear artifact preparation list
- Embedded “proactive” behavioral probes
- Adaptive difficulty options
  

---

Artifacts to Prepare Before Interview

1. Git Repository (Repo A)
  - Branches: main, feature/broken-ci
  - Intentionally flawed .gitlab-ci.yml
  - Go service (unchecked error + hardcoded placeholder credential)
  - Optional Python utility (e.g., unsafe subprocess use)
2. Pre-run Scanner Output
  - Gosec report: 2 true positives, 2 likely false positives
  - TruffleHog: 1 real-looking secret, 1 false positive
3. Pipeline Failure Log (missing dependency / bad stage ordering)
4. Code Review Snippet
5. Kubernetes Manifest (misconfigs: privileged, no limits, hardcoded secret)
6. Optional SBOM (CycloneDX/Syft) for advanced SCA prioritization discussion
  

---

Segment Details & Prompts
1. Intro (5 min)
- “Which SSDLC phase is most neglected and how would you improve it?”
  
2. Git & Repo Handling (8 min)
- Clone & inspect branches; locate CI YAML issues.
- “Local vs server-side hooks—enforcement differences?”
- “Risks: shared vs private runners?”
  
3. Scanner Result Interpretation (12 min)
- Categorize findings (true vs false positive).
- “Reduce secret scan noise without suppressing real leaks?”
- “Prioritization across SCA / SAST / Secrets?”
  
4. Pipeline Extension & Debug (15 min)
- Verbally add SCA (DependencyTrack/Snyk) + optional IaC scanner (Checkov).
- Fix YAML: staging order, secret handling, caching, runner isolation.
- “How would you secure third-party action/template usage?”
  
5. Code Review & Secure Fix (10 min)
- Identify vulnerabilities (path traversal, temp file race, unchecked errors).
- Propose secure refactor + logging strategy.
  
6. Kubernetes & SSDLC (5 min)
- Rank misconfig impacts (privilege, secret exposure, lack of limits).
- “Automate baseline policy enforcement?”
  
7. Incident & Zero-Day Response (5 min)
- “First 60 minutes after disclosure?”
- “Transitive dependency exposure mapping approach?”
  
8. Wrap-Up (≈5 min)
- Candidate questions; assess curiosity, improvement mindset.
  

---

Expanded Question Bank (Selected)
Security Scanners:
- SAST vs DAST vs IAST vs RASP differences
- “Reachability analysis—why useful for SCA prioritization?”
  
CI/CD:
- “Design merge gating strategy (fast-lint, secrets, SAST pre-merge; SCA post-merge).”
- “Policy-as-code examples (OPA/Gatekeeper, Conftest).”
  
Git:
- “Enforce signed commits—why does it matter?”
- “Pre-receive hook to block leaked secrets—outline logic.”
  
Incident Response:
- “Zero-day in transitive dependency—tooling path (SBOM graph).”
  
Kubernetes:
- “Runtime hardening beyond image scanning (Seccomp, AppArmor, eBPF telemetry).”
  
Proactivity & Culture:
- “Security metrics to track improvement (MTTR for critical vulns, false positive ratio trend).”
  
AI Usage:
- “Controls to verify AI-generated remediation suggestions.”
  

---
Evaluation Rubric (Weighted Scoring)

Scoring Model:
- Three tiers per category: Below Expectation (B), Meets Expectation (M), Above Expectation (A)
- Suggested numeric mapping: B = 0, M = 1, A = 2
- Category Score = Tier Value × Weight
- Max composite score = Sum(Weight × 2)
- Decision guidance (example thresholds; adjust per hiring bar):
  - Strong Hire: ≥75% of max AND ≥5 categories at A
  - Hire: ≥55% AND no critical category (Pipeline, Secure Coding, Incident Response) at B
  - Lean / No Hire: <55% OR ≥2 critical categories at B
    
Category
Weight
Below Expectation (0)
Meets Expectation (1)
Above Expectation (2)
Pipeline Architecture & Hardening
20%
Struggles to fix basic YAML; no concept of stage optimization; ignores secrets handling
Corrects syntax, orders scanners logically; basic secret masking; mentions runner isolation superficially
Designs minimal-fail path, parallelization, ephemeral runners, OIDC / short‑lived creds, caching security; articulates policy gates & supply-chain integrity
Vulnerability Triage & Prioritization
20%
Relies solely on CVSS; mislabels false positives; no prioritization model
Distinguishes false positives; considers exploit maturity & asset criticality
Provides multi-factor model (reachability, blast radius, compensating controls, business impact); suggests automation workflow & SLA tracking
Secure Coding & Code Review
20%
Misses obvious vulnerabilities; offers superficial fixes
Identifies main issues; proposes reasonable refactor; addresses input validation & error handling
Comprehensive root-cause reasoning; secure patterns (least privilege, immutability); considers performance trade-offs & adds tests/documentation
Git Mastery & Workflow Governance
20%
Basic clone/commit only; unclear on hooks or branching strategy
Explains hooks, branching model, signed commits concept
Designs enforcement (server hooks, commit signing, semantic versioning gating, protected branches, pre-receive policy integration)
SSDLC Integration & Shift-Left Strategy
10%
Sees SSDLC as a list of scans; no lifecycle linkage
Maps scanners to phases; mentions training & threat modeling
Integrates feedback loops, risk-based gates, metrics at each phase; continuous improvement plan
Kubernetes Security Awareness
10%
Identifies few misconfigs; suggests generic “add limits”
Spots key misconfigs (privileged user, secrets); proposes baseline policies
Prioritizes by exploit path; suggests admission controls, runtime detection, segregation, Pod Security Standards, workload identity

Pass / Concern Indicators (Refined)
Pass Indicators:
- Provides structured, prioritized answers (risk-based)
- Balances automation vs developer experience
- Demonstrates ownership mindset (suggests improvements proactively)
- Connects tooling to lifecycle outcomes (MTTR, quality gates)
  
Concern Indicators:
- Tool-name recitation without mechanics
- No differentiation between severity and exploitability
- Ignores secrets & runner isolation in pipeline changes
- Fails to recognize high-impact K8 misconfigs

Example Expected Outputs Per Category
Pipeline (M):
- “Add stages: validate -> test -> scan -> build -> deploy; secrets scan early to fail fast; pin images; use masked CI variables.”
  
Pipeline (A):
- Same as above plus: ephemeral isolated runner, concurrency limits, parallel SAST/SCA after compile caching, SBOM generation step
  
Vulnerability Triage (M):
- Categorizes findings, deprioritizes low exploitability medium severity, focuses on high severity reachable code paths.
  
Vulnerability Triage (A):
- Provides scoring matrix: (Reachability + Exploit Maturity + Asset Criticality + Exposure Surface); suggests automation pipeline labeling & dashboard.
  
(Proctors can record short exemplar outputs under each rubric row.)

---
Recommended Guideline Enhancements (Retained)

1. Standardize environment & synthetic data.
2. Timeboxing with visible timer; adaptive depth.
3. Clarify no need for memorizing CLI flags.
4. Add optional supply chain / provenance if candidate excels.
5. Structured note capture immediately post-interview.
6. Optional asynchronous follow-up design challenge.
  

---

Sample Broken .gitlab-ci.yml

Common issues: missing stages, job order misalignment, unpinned image, secrets echoed, ambiguous branch rules.
image: golang:latest

test_job:
  stage: tests
  script:
    - go test ./...
    - env > version.txt
  only:
    - main
  artifacts:
      path:
          - version.txt

sca_scan:
  script:
    - snyk test
  dependencies:
    - test_job

deploy_prod:
  stage: deploy
  script:
    - echo "DEPLOY_TOKEN=$DEPLOY_TOKEN"
    - ./deploy.sh

Expected Fix (Meets):
- Add stages
- Pin image version
- Remove secret echo
- Correct stage names
- Introduce branch/MR rules
- Add SCA job
- Fix artifact expiry
  
Expected Fix (Above):
- Add validate (secret scan) stage


---
Secure Refactor Talking Points (Compress Script Example)
Issues:
- Unsanitized path
- Insecure temp file handling
- Unbounded resource usage
- Missing error propagation
- Potential race conditions
  
Mitigations:
- filepath.Clean, validate root scope
- os.CreateTemp with restrictive perms
- Centralized structured logging + error wrapping
- Add unit test for path traversal attempt
  

---

Optional Extensions
- Evolution plan: baseline -> metrics-driven maturity -> formal SLAs
  

---
Closing
End with candidate questions; probe “What would you implement in your first 90 days?” to assess strategic and proactive alignment

---
