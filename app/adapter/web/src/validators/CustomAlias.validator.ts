export function validateCustomAliasFormat(customAlias?: string): string | null {
    const CUSTOM_ALIAS_MAX_LENGTH = 50;

    if (!customAlias) {
        return null;
    }

    if (customAlias.length >= CUSTOM_ALIAS_MAX_LENGTH) {
        return `Custom alias has to contain less than ${CUSTOM_ALIAS_MAX_LENGTH} characters`;
    }

    return null;
}