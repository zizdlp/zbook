"use client";
import { Popover, Transition } from "@headlessui/react";
import { Fragment } from "react";
import { MdOutlineMoreVert } from "react-icons/md";
import DeleteCommentDialog from "./DeleteCommentDialog";
import React, { useState, useContext } from "react";
import { OperationContext } from "@/providers/OperationProvider";
import { toast } from "react-toastify";
import { CommentCountInfo } from "@/types/model";
import { useTranslations } from "next-intl";
interface CommentOperationMoreProps {
  comment_id: number;
  markdown_id: number;
  setIsDeleted: React.Dispatch<React.SetStateAction<boolean>>;
  commentCountInfo: CommentCountInfo | undefined;
  owned: boolean;
}
export default function CommentOperationMore(props: CommentOperationMoreProps) {
  const t = useTranslations("Dialog");
  const [showDialog, setShowDialog] = useState(false);
  const { setCreateCommentReportOpen, setOperationCommentID } =
    useContext(OperationContext);
  return (
    <div className="">
      <DeleteCommentDialog
        showDialog={showDialog}
        setShowDialog={setShowDialog}
        comment_id={props.comment_id}
        markdown_id={props.markdown_id}
        setIsDeleted={props.setIsDeleted}
      />
      <Popover>
        {({ open }) => (
          <>
            <Popover.Button>
              <MdOutlineMoreVert />
            </Popover.Button>
            <Transition
              as={Fragment}
              enter="transition ease-out duration-200"
              enterFrom="opacity-0 "
              enterTo="opacity-100 "
              leave="transition ease-in duration-150"
              leaveFrom="opacity-100 "
              leaveTo="opacity-0 "
            >
              <Popover.Panel
                className="border border-slate-300 dark:border-slate-700 absolute cursor-pointer z-50 rounded-md  -translate-y-10 -translate-x-20
                bg-slate-100 dark:bg-slate-800/75 text-slate-700 dark:text-slate-200 font-semibold shadow-lg dark:backdrop-blur-md"
              >
                {props.owned && (
                  <div
                    onClick={() => {
                      setShowDialog(true);
                    }}
                    className="flex items-center rounded-t-md justify-center hover:dark:bg-sky-900 hover:bg-sky-200 px-4 py-2 min-w-[4rem]"
                  >
                    <span>{t("Delete")}</span>
                  </div>
                )}
                {!props.owned && (
                  <div
                    onClick={() => {
                      if (props.commentCountInfo?.is_reported) {
                        toast(t("AlreadyReported"), {
                          type: "error",
                          isLoading: false,
                          autoClose: 1500,
                        });
                        return;
                      }
                      setOperationCommentID(props.comment_id);
                      setCreateCommentReportOpen(true);
                    }}
                    className="flex items-center rounded-t-md justify-center hover:dark:bg-sky-900 hover:bg-sky-200 px-4 py-2 min-w-[4rem]"
                  >
                    {t("Report")}
                  </div>
                )}
              </Popover.Panel>
            </Transition>
          </>
        )}
      </Popover>
    </div>
  );
}
