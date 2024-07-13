"use client";

import React, { useState } from "react";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { useTranslations } from "next-intl";
export default function UpdateCommentReportButton({
  report_id,
  processed,
}: {
  report_id: number;
  processed: boolean;
}) {
  const [isProcessed, setIsProcessed] = useState(processed);

  async function updateCommentReport() {
    fetchServerWithAuthWrapper({
      endpoint: FetchServerWithAuthWrapperEndPoint.UPDATE_COMMENT_REPORT_STATUS,
      xforward: "",
      agent: "",
      tags: [],
      values: {
        report_id: report_id,
        processed: !isProcessed,
      },
    }).then((data) => {
      setIsProcessed(!isProcessed);
    });
  }
  const t = useTranslations("DataList");
  return (
    <div
      className="bg-sky-500 dark:bg-sky-700 text-white rounded-lg py-2 px-4 font-semibold text-sm cursor-pointer hover:bg-sky-600 dark:hover:bg-sky-500"
      onClick={() => updateCommentReport()}
    >
      <span className="flex-1 whitespace-nowrap">
        {isProcessed ? t("Resolved") : t("Unresolved")}
      </span>
    </div>
  );
}
