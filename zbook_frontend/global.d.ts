import en from "./messages/en.json";
import zh from "./messages/zh.json";

// Ensure all language files conform to the Messages type
type EnMessages = typeof en;
type ZhMessages = typeof zh;

declare global {
  // Use type safe message keys with `next-intl`
  interface IntlMessages extends ZhMessages, EnMessages {}
}
