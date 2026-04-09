# 6장. Enterprise를 위한 기반 정비

> **스토리**: 고객이 늘면서 느려지고, 보안도 허술하다 → 대형 고객사를 받기 전에 기반을 정비하자

Valkey 캐시로 성능을 개선하고, Google Secret Manager로 시크릿을 안전하게 관리합니다. Blue/Green에서 Canary 배포로 전환하여 더 안전한 배포 전략을 적용합니다.

---

| 절 | 제목 |
|:---:|------|
| 6.1 | <small>Valkey 캐시: 느려진 서비스 개선</small> |
| 6.2 | <small>Secret 관리: Google Secret Manager + Secrets Store CSI Driver</small> |
| 6.3 | <small>Canary 배포: B/G보다 신중하게, 점진적으로</small> |
| 💡 | <small>Hints & Tips: `claude-context/`로 현재 아키텍처 정리하기</small> |

> 자세한 내용은 책을 참고하세요.
