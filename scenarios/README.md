# DevSecOps Senior Interview Scenarios

This directory contains pre-configured scenarios and artifacts for the 60-minute structured interview.

## Scenario Overview

### Git Repository Structure
- **main branch**: Contains working application with vulnerabilities
- **feature/broken-ci branch**: Contains broken CI/CD pipeline for debugging exercise
- **Commit history**: Includes commits with exposed secrets for Git handling assessment

### Scanner Output Analysis
- **gosec-results.json**: Contains 6 findings (4 true positives, 2 potential false positives)
- **trufflehog-results.json**: Contains 3 secret detection results (2 real secrets, 1 false positive)

### Pipeline Scenarios  
- **broken-pipeline.yml**: GitHub Actions workflow with multiple security issues
- **.gitlab-ci.yml**: GitLab CI with staging problems and secret exposure
- **pipeline-failure.log**: Sample failure log for debugging exercise

### Code Review Targets
- **cmd/main.go**: Hardcoded credentials, missing error handling
- **pkg/compress/compress.go**: Unbounded resource usage, temp file race conditions
- **pkg/fileutils/fileutils.go**: Insecure permissions, path validation bypass
- **internal/handlers/handlers.go**: Path traversal, command injection

### Infrastructure Misconfigurations
- **k8s/deployment.yaml**: Privileged containers, host access, missing limits
- **terraform/main.tf**: Open security groups, unencrypted storage

## Interview Flow Usage

### Phase 1: Git & Repository (8 min)
1. Have candidate clone repository
2. Explore branch structure and identify CI issues
3. Discuss Git security practices (hooks, signed commits)

### Phase 2: Scanner Results (12 min)  
1. Present gosec-results.json and trufflehog-results.json
2. Ask candidate to triage findings (true vs false positives)
3. Discuss prioritization strategy and noise reduction

### Phase 3: Pipeline Extension (15 min)
1. Show broken pipeline configurations
2. Ask to fix YAML issues and add SCA/IaC scanning
3. Discuss secure CI/CD practices

### Phase 4: Code Review (10 min)
1. Focus on specific vulnerable functions
2. Ask to identify security issues and propose fixes
3. Evaluate secure coding knowledge

### Phase 5: Kubernetes Security (5 min)
1. Review deployment.yaml misconfigurations  
2. Prioritize security issues by impact
3. Discuss policy enforcement approaches

### Phase 6: Incident Response (5 min)
1. Present Log4Shell-style scenario
2. Discuss first 60 minutes response plan
3. Evaluate dependency management approach

## Expected Findings Summary

### High Priority Issues
- Hardcoded API key in source code (G101)
- Privileged Kubernetes containers  
- Command injection vulnerability
- Path traversal in file serving
- Host filesystem access in containers

### Medium Priority Issues  
- Insecure file permissions (G302)
- Missing resource limits
- Unhandled errors (G104)
- Secret exposure in CI logs

### Triage Challenges
- Test file secrets (false positive)
- Example credentials in documentation
- Development vs production context

## Scoring Guidance

Refer to INTERVIEW_GUIDE.md for detailed rubric, but key indicators:

**Strong Performance:**
- Systematically categorizes findings by exploitability
- Suggests automation for triage workflow
- Designs comprehensive fix strategies
- Demonstrates defense-in-depth thinking

**Concerns:**  
- Cannot distinguish false positives
- Focuses only on CVSS scores
- Misses critical Kubernetes misconfigurations
- No consideration of developer experience