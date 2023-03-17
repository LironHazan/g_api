const { composePlugins, withNx } = require('@nrwl/webpack');
const { withReact } = require('@nrwl/react');
const { withModuleFederation } = require('@nrwl/react/module-federation');
const { merge } = require('webpack-merge');
// const { dependencies } = require('../../package.json');

// Nx plugins for webpack.
module.exports = composePlugins(withNx(), withReact(), async (config, { options, context }) => {
  // Update the webpack config as needed here.
  // e.g. config.plugins.push(new MyPlugin())
  // For more information on webpack config and Nx see:
  // https://nx.dev/packages/webpack/documents/webpack-config-setup
  const federatedModules = await withModuleFederation({
    name: 'repro-issue-15475-remote',
    exposes: {
      './Module': './src/remote-entry.ts',
    },
    shared: (dep, sconfig) => {
      if (dep.startsWith('@apollo/client')) {
        return false;
      }
      return {
        ...sconfig,
        eager: true,
        strictVersion: false,
      };
    }
  });

  return merge(federatedModules(config, context), {
    // overwrite values here
  });
});
