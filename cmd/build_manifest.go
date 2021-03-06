package cmd

import (
	boshtpl "github.com/cloudfoundry/bosh-cli/director/template"
	boshui "github.com/cloudfoundry/bosh-cli/ui"
)

type BuildManifestCmd struct {
	ui boshui.UI
}

func NewBuildManifestCmd(ui boshui.UI) BuildManifestCmd {
	return BuildManifestCmd{ui: ui}
}

func (c BuildManifestCmd) Run(opts BuildManifestOpts) error {
	tpl := boshtpl.NewTemplate(opts.Args.Manifest.Bytes)

	vars := opts.VarFlags.AsVariables()
	ops := opts.OpsFlags.AsOps()
	evalOpts := boshtpl.EvaluateOpts{ExpectAllKeys: opts.VarErrors}

	bytes, err := tpl.Evaluate(vars, ops, evalOpts)
	if err != nil {
		return err
	}

	c.ui.PrintBlock(string(bytes))

	return nil
}
