package service

type RiskResult string

const (
    REJECTED        RiskResult = "REJECTED"
    APPROVED_LOW    RiskResult = "APPROVED_LOW"
    PENDING_REVIEW  RiskResult = "PENDING_REVIEW"
)

func AssessRisk(creditScore int) RiskResult {
    if creditScore < 600 {
        return REJECTED
    }
    if creditScore >= 720 {
        return APPROVED_LOW
    }
    return PENDING_REVIEW
}
