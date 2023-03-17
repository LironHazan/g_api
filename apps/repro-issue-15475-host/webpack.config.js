const { composePlugins, withNx } = require('@nrwl/webpack');
const { merge } = require('webpack-merge');
const { withReact } = require('@nrwl/react');
const { withModuleFederation } = require('@nrwl/react/module-federation');
const { dependencies } = require('../../package.json');

// Nx plugins for webpack.
module.exports = composePlugins(
  withNx(),
  withReact(),
  async (config, { options, context }) => {
    // Update the webpack config as needed here.
    // e.g. config.plugins.push(new MyPlugin())
    // For more information on webpack config and Nx see:
    // https://nx.dev/packages/webpack/documents/webpack-config-setup
    const federatedModules = await withModuleFederation({
      name: 'repro-issue-15475-host',
      remotes: [
        ['repro-issue-15475-remote', '//localhost:3000/'],
      ],
      // Tod avoid: Uncaught Error: Shared module is not available for eager consumption,
      // Should specify the shared function like in the tests here --> https://github.com/nrwl/nx/pull/10156/files
      shared: (dep, sconfig) => {
        if (dep.startsWith('@apollo/client')) {
          //return false;
          return {
            eager: true,
            requiredVersion: dependencies['@apollo/client'],
            strictVersion: true,
          };
        }
        return {
          ...sconfig,
          eager: true,
          strictVersion: false,
        };
      },
      additionalShared: [
        '@apollo/client',
        'graphql',
      ]
    });

    return merge(federatedModules(config, context), {
      // overwrite values here
    });

    // return config;
});
