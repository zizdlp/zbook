"use client";

import React, { createContext, useState } from "react";
export const OperationContext = createContext<{
  createCommentOpen: boolean; // open create comment dialog
  setCreateCommentOpen: React.Dispatch<React.SetStateAction<boolean>>;

  deleteCommentOpen: boolean; // open create comment dialog
  setDeleteCommentOpen: React.Dispatch<React.SetStateAction<boolean>>;

  updateUserOpen: boolean; // open create comment dialog
  setUpdateUserOpen: React.Dispatch<React.SetStateAction<boolean>>;
  deleteUserOpen: boolean; // open create comment dialog
  setDeleteUserOpen: React.Dispatch<React.SetStateAction<boolean>>;

  showVisibleOpen: boolean; // open create comment dialog
  setShowVisibleOpen: React.Dispatch<React.SetStateAction<boolean>>;

  updateRepoOpen: boolean; // open create comment dialog
  setUpdateRepoOpen: React.Dispatch<React.SetStateAction<boolean>>;
  deleteRepoOpen: boolean; // open create comment dialog
  setDeleteRepoOpen: React.Dispatch<React.SetStateAction<boolean>>;

  createCommentReportOpen: boolean; // open create comment report dialog
  setCreateCommentReportOpen: React.Dispatch<React.SetStateAction<boolean>>;

  createRepoReportOpen: boolean; // open create post dialog
  setCreateRepoReportOpen: React.Dispatch<React.SetStateAction<boolean>>;

  CreateSystemNotificationOpen: boolean;
  setCreateSystemNotificationOpen: React.Dispatch<
    React.SetStateAction<boolean>
  >;

  createInvitationOpen: boolean;
  setCreateInvitationOpen: React.Dispatch<React.SetStateAction<boolean>>;

  createRepoOpen: boolean; // open create comment dialog
  setCreateRepoOpen: React.Dispatch<React.SetStateAction<boolean>>;

  mutationCreateComment: boolean; // create comment
  setMutationCreateComment: React.Dispatch<React.SetStateAction<boolean>>;

  createCommentContent: string; // create comment
  setCreateCommentContent: React.Dispatch<React.SetStateAction<string>>;

  mutationCreateRepo: boolean; // create comment
  setMutationCreateRepo: React.Dispatch<React.SetStateAction<boolean>>;

  mutationDeleteComment: boolean; // delete comment
  setMutationDeleteComment: React.Dispatch<React.SetStateAction<boolean>>;
  mutationUpdateComment: boolean; // update comment
  setMutationUpdateComment: React.Dispatch<React.SetStateAction<boolean>>;

  operationRepoName: string;
  setOperationRepoName: React.Dispatch<React.SetStateAction<string>>;

  operationMarkdownID: number;
  setOperationMarkdownID: React.Dispatch<React.SetStateAction<number>>;

  operationUsername: string;
  setOperationUsername: React.Dispatch<React.SetStateAction<string>>;

  operationCommentID: number;
  setOperationCommentID: React.Dispatch<React.SetStateAction<number>>;
  operationRootID: number;
  setOperationRootID: React.Dispatch<React.SetStateAction<number>>;
  operationParentID: number;
  setOperationParentID: React.Dispatch<React.SetStateAction<number>>;

  mutationToggleTheme: boolean; // create comment
  setMutationToggleTheme: React.Dispatch<React.SetStateAction<boolean>>;
}>({
  createCommentOpen: false,
  setCreateCommentOpen: () => {},
  createInvitationOpen: false,
  setCreateInvitationOpen: () => {},
  deleteCommentOpen: false,
  setDeleteCommentOpen: () => {},

  updateUserOpen: false,
  setUpdateUserOpen: () => {},
  deleteUserOpen: false,
  setDeleteUserOpen: () => {},

  showVisibleOpen: false,
  setShowVisibleOpen: () => {},

  updateRepoOpen: false,
  setUpdateRepoOpen: () => {},
  deleteRepoOpen: false,
  setDeleteRepoOpen: () => {},

  createCommentReportOpen: false,
  setCreateCommentReportOpen: () => {},

  createRepoOpen: false,
  setCreateRepoOpen: () => {},

  createRepoReportOpen: false,
  setCreateRepoReportOpen: () => {},

  mutationCreateComment: false,
  setMutationCreateComment: () => {},
  createCommentContent: "",
  setCreateCommentContent: () => {},

  mutationCreateRepo: false,
  setMutationCreateRepo: () => {},

  mutationUpdateComment: false,
  setMutationUpdateComment: () => {},
  mutationDeleteComment: false,
  setMutationDeleteComment: () => {},

  operationRepoName: "",
  setOperationRepoName: () => {},

  operationMarkdownID: 0,
  setOperationMarkdownID: () => {},

  operationUsername: "",
  setOperationUsername: () => {},

  operationCommentID: 0,
  setOperationCommentID: () => {},
  operationParentID: 0,
  setOperationParentID: () => {},
  operationRootID: 0,
  setOperationRootID: () => {},

  CreateSystemNotificationOpen: false,
  setCreateSystemNotificationOpen: () => {},
  mutationToggleTheme: false,
  setMutationToggleTheme: () => {},
});

export default function OperationProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [createCommentOpen, setCreateCommentOpen] = useState(false);
  const [createInvitationOpen, setCreateInvitationOpen] = useState(false);
  const [deleteCommentOpen, setDeleteCommentOpen] = useState(false);
  const [updateUserOpen, setUpdateUserOpen] = useState(false);
  const [deleteUserOpen, setDeleteUserOpen] = useState(false);
  const [showVisibleOpen, setShowVisibleOpen] = useState(false);

  const [updateRepoOpen, setUpdateRepoOpen] = useState(false);
  const [deleteRepoOpen, setDeleteRepoOpen] = useState(false);
  const [createRepoOpen, setCreateRepoOpen] = useState(false);

  const [createCommentReportOpen, setCreateCommentReportOpen] = useState(false);

  const [createRepoReportOpen, setCreateRepoReportOpen] = useState(false);
  const [mutationCreateComment, setMutationCreateComment] = useState(false);
  const [createCommentContent, setCreateCommentContent] = useState("");

  const [mutationCreateRepo, setMutationCreateRepo] = useState(false);

  const [mutationUpdateComment, setMutationUpdateComment] = useState(false);
  const [mutationDeleteComment, setMutationDeleteComment] = useState(false);

  const [operationUsername, setOperationUsername] = useState("");
  const [operationRepoName, setOperationRepoName] = useState("");
  const [operationMarkdownID, setOperationMarkdownID] = useState(0);
  const [operationParentID, setOperationParentID] = useState(0);
  const [operationCommentID, setOperationCommentID] = useState(0);
  const [operationRootID, setOperationRootID] = useState(0);
  const [mutationToggleTheme, setMutationToggleTheme] = useState(false);
  const [CreateSystemNotificationOpen, setCreateSystemNotificationOpen] =
    useState(false);
  // Context values passed to consumer
  const value = {
    createCommentOpen,
    setCreateCommentOpen,
    createInvitationOpen,
    setCreateInvitationOpen,
    deleteCommentOpen,
    setDeleteCommentOpen,
    createRepoOpen,
    setCreateRepoOpen,
    updateUserOpen,
    setUpdateUserOpen,
    deleteUserOpen,
    setDeleteUserOpen,
    showVisibleOpen,
    setShowVisibleOpen,

    updateRepoOpen,
    setUpdateRepoOpen,
    deleteRepoOpen,
    setDeleteRepoOpen,

    createCommentReportOpen,
    setCreateCommentReportOpen,

    createRepoReportOpen,
    setCreateRepoReportOpen,
    mutationCreateComment,
    setMutationCreateComment,
    createCommentContent,
    setCreateCommentContent,

    mutationCreateRepo,
    setMutationCreateRepo,

    mutationUpdateComment,
    setMutationUpdateComment,
    mutationDeleteComment,
    setMutationDeleteComment,

    operationRepoName,
    setOperationRepoName,
    operationUsername,
    setOperationUsername,
    operationMarkdownID,
    setOperationMarkdownID,

    operationParentID,
    setOperationParentID,
    operationCommentID,
    setOperationCommentID,
    operationRootID,
    setOperationRootID,
    CreateSystemNotificationOpen,
    setCreateSystemNotificationOpen,
    mutationToggleTheme,
    setMutationToggleTheme,
  };
  return (
    <OperationContext.Provider value={value}>
      {children}
    </OperationContext.Provider>
  );
}
