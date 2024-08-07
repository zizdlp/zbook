import "../globals.css";
import "react-toastify/dist/ReactToastify.css";
import SearchDialog from "@/providers/dialogs/SearchDialog";
import SideBarProvider from "@/providers/SideBarProvider";
import SearchDialogProvider from "@/providers/SearchDialogProvider";
import { ThemeProvider } from "next-themes";
import NotiDialogProvider from "@/providers/NotiDialogProvider";
import { ToastContainer } from "react-toastify";
import NavBar from "@/components/navbars/NavBar";
import NotificationDialog from "@/components/notifications/NotificationDialog";
import OperationProvider from "@/providers/OperationProvider";
import CreateCommentReportDialog from "@/providers/dialogs/CreateCommentReportDialog";
import CreateSystemNotification from "@/providers/dialogs/CreateSystemNotification";

import UpdateUserDialog from "@/providers/dialogs/UpdateUserDialog";

import CreateRepoDialog from "@/providers/dialogs/CreateRepoDialog";
import UpdateRepoDialog from "@/providers/dialogs/UpdateRepoDialog";
import DeleteRepoDialog from "@/providers/dialogs/DeleteRepoDialog";
import { Metadata } from "next";
import { NextIntlClientProvider } from "next-intl";
import { ReactNode } from "react";
import LogVisitor from "@/components/LogVisitor";
import DeleteUserDialog from "@/providers/dialogs/DeleteUserDialog";
import { SessionProvider } from "next-auth/react";
import CreateCommentDialog from "@/components/comments/CreateCommentDialog";
import {
  getMessages,
  getTranslations,
  unstable_setRequestLocale,
} from "next-intl/server";

export async function generateMetadata({
  params: { locale },
}: Omit<Props, "children">): Promise<Metadata> {
  const t = await getTranslations({ locale, namespace: "GenerateMetaData" });
  return {
    title: {
      template: "%s - ZBook",
      default: "ZBook | " + t("Slogan"), // a default is required when creating a template
    },
  };
}

type Props = {
  children: ReactNode;
  params: { locale: string };
};
import { JetBrains_Mono } from "next/font/google";
import CreateInvitation from "@/providers/dialogs/CreateInvitation";

const jetbrains_mono = JetBrains_Mono({
  subsets: ["latin"],
  display: "swap",
  variable: "--font-jetbrains-mono",
});

export default async function LocaleLayout({
  children,
  params: { locale },
}: Props) {
  // Enable static rendering
  unstable_setRequestLocale(locale);
  const messages = await getMessages();
  return (
    <html
      lang={locale}
      className={`${jetbrains_mono.variable} `}
      suppressHydrationWarning
    >
      <head>
        <link rel="icon" href="/favicon.ico" sizes="any" />
        <link
          rel="apple-touch-icon"
          href="/logo_256.png"
          type="image/png"
          sizes="256x256"
        />
      </head>
      <body className="antialiased text-slate-700 dark:text-slate-200 bg-white dark:bg-slate-900 min-h-screen">
        <NextIntlClientProvider messages={messages}>
          <SessionProvider>
            <SearchDialogProvider>
              <NotiDialogProvider>
                <OperationProvider>
                  <ThemeProvider attribute="class">
                    <SideBarProvider>
                      <NavBar />
                      <NotificationDialog />
                      <CreateSystemNotification />
                      <CreateInvitation />
                      <CreateCommentReportDialog />
                      <CreateCommentDialog />
                      <CreateRepoDialog />
                      <UpdateRepoDialog />
                      <DeleteRepoDialog />
                      <UpdateUserDialog />
                      <DeleteUserDialog />
                      <ToastContainer position="bottom-right" theme="dark" />
                      <SearchDialog />
                      <LogVisitor />
                      {children}
                    </SideBarProvider>
                  </ThemeProvider>
                </OperationProvider>
              </NotiDialogProvider>
            </SearchDialogProvider>
          </SessionProvider>
        </NextIntlClientProvider>
      </body>
    </html>
  );
}
