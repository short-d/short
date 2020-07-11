module.exports = {
    stories: ['../src/**/*.stories.tsx'],
    addons: [
        '@storybook/addon-knobs/register',
        '@storybook/addon-actions/register'
    ],
    webpackFinal: async config => {
        config.module.rules.push({
            test: /\.(ts|tsx)$/,
            use: ['ts-loader', 'react-docgen-typescript-loader'],
        }, {
            test: /\.scss$/,
            use: [
                'style-loader',
                {
                    loader: 'css-loader',
                    options: {
                        modules: true,
                    },
                },
                'sass-loader'],
        });
        config.resolve.extensions.push('.ts', '.tsx', '.scss');
        return config;
    },
};