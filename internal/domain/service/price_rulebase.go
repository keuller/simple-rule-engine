package service

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/keuller/pricing-service/internal/domain/repository"
)

var (
	version int = 0
)

type PriceRuleBase struct {
	dataCtx       ast.IDataContext
	knowledgeBase *ast.KnowledgeLibrary
	repo          repository.PriceRepository
}

func NewPriceRuleBase(repo repository.PriceRepository) PriceRuleBase {
	ruleCtx := ast.NewDataContext()
	prb := PriceRuleBase{ruleCtx, nil, repo}
	prb.BuildRuleBase()
	return prb
}

func (prb *PriceRuleBase) AddFact(name string, fact interface{}) error {
	if err := prb.dataCtx.Add(name, fact); err != nil {
		return err
	}
	return nil
}

func (prb *PriceRuleBase) BuildRuleBase() {
	version++
	rules := prb.repo.GetRules()

	var output strings.Builder
	prb.knowledgeBase = ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(prb.knowledgeBase)

	tmpl, _ := template.New("rules").Parse(`
	rule {{.ID}} "{{.Title}}" {
		when
			{{.When}}
		then
			{{.Then}}
	}`)

	for idx := range rules {
		buf := bytes.NewBufferString("")
		_ = tmpl.Execute(buf, rules[idx])
		output.WriteString(buf.String())
	}

	// log.Println("Rules:", output.String())

	baseVersion := fmt.Sprintf("%d.0", version)
	log.Println("Rule base version:", baseVersion)
	bs := pkg.NewBytesResource([]byte(output.String()))
	if err := ruleBuilder.BuildRuleFromResource("DPRules", baseVersion, bs); err != nil {
		log.Printf("[ERROR] Fail to build base rules - %v \n", err)
	}
}

func (prb *PriceRuleBase) Process() {
	baseVersion := fmt.Sprintf("%d.0", version)
	log.Println("Processing rule base version:", baseVersion)
	instance := prb.knowledgeBase.NewKnowledgeBaseInstance("DPRules", baseVersion)
	engine := engine.NewGruleEngine()
	engine.MaxCycle = 10000
	if err := engine.Execute(prb.dataCtx, instance); err != nil {
		log.Printf("[ERROR] Fail to process rules - %v \n", err)
	}

	prb.dataCtx = ast.NewDataContext()
}
