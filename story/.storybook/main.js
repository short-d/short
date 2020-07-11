module.exports = {
    stories: ['../src/**/*.stories.tsx'],
    webpackFinal: async config => {
        config.module.rules.push({
            test: /\.(ts|tsx)$/,
            use: ['ts-loader'],
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