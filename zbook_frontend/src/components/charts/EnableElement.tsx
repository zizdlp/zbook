import { useState } from "react";
import { Switch } from "@headlessui/react";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import ToolTip from "@/components/ToolTip";
import { useTranslations } from "next-intl";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";

export default function EnableElement({
  config_name,
  label,
  initEnabled,
}: {
  config_name: string;
  label: string;
  initEnabled: boolean;
}) {
  const [enabled, setEnabled] = useState(initEnabled);
  const t = useTranslations("AdminOverView");
  return (
    <Switch.Group>
      <div className="flex items-center">
        <Switch.Label className="mr-1">
          <ToolTip
            message={
              config_name == "allow_login"
                ? t("TipLogin")
                : config_name == "allow_register"
                  ? t("TipRegister")
                  : t("TipInviation")
            }
          >
            <div>{label}</div>
          </ToolTip>
        </Switch.Label>
        <Switch
          checked={enabled}
          onChange={async () => {
            try {
              const data = await fetchServerWithAuthWrapper({
                endpoint:
                  FetchServerWithAuthWrapperEndPoint.UpdateConfiguration,
                xforward: "",
                agent: "",
                tags: [],
                values: { config_name: config_name, config_value: !enabled },
              });
              if (data.error) {
                throw new FetchError(data.message, data.status);
              }
              setEnabled(!enabled);
            } catch (error) {
              let e = error as FetchError;
              logger.error(`setEnabled failed:${e.message}`, e.status);
            }
          }}
          className={`${
            enabled ? "bg-green-600" : "bg-gray-200 dark:bg-gray-700"
          } relative inline-flex h-6 w-11 items-center rounded-full transition-colors focus:outline-none focus:ring-none`}
        >
          <span
            className={`${
              enabled ? "translate-x-6" : "translate-x-1"
            } inline-block h-4 w-4 transform rounded-full bg-white transition-transform`}
          />
        </Switch>
      </div>
    </Switch.Group>
  );
}
