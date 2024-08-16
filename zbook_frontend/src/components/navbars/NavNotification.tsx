"use client";
import { useEffect, useContext, useRef } from "react";
import { Badge } from "antd";
import { MdOutlineNotificationsNone } from "react-icons/md";
import { NotiDialogContext } from "@/providers/NotiDialogProvider";
import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { FetchError } from "@/fetchs/util";
import { logger } from "@/utils/logger";
import { toast } from "react-toastify";
import { useTranslations } from "next-intl";
import { joinUrlParts } from "@/utils/util";

export default function NavNotification({
  username,
  websocket_url,
  unread_count,
  access_token,
}: {
  username: string;
  websocket_url: string;
  unread_count: number;
  access_token: string;
}) {
  const t = useTranslations("Toast");
  const {
    notiDialogOpen,
    setNotiDialogOpen,
    mutationReadNotification,
    setMutationReadNotification,
    unReadCount,
    setUnReadCount,
  } = useContext(NotiDialogContext);
  const reconnectRef = useRef(0); // 使用useRef来保存reconnect的引用
  async function resetUnreadCount() {
    try {
      const data = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.RESET_UNREAD_COUNT,
        tags: [],
        xforward: "",
        agent: "",
        values: {},
      });
      if (data.error) {
        throw new FetchError(data.message, data.status);
      }
    } catch (error) {
      let e = error as FetchError;
      logger.error(`reset unreadcount failed:${e.message}`, e.status);
    }
  }
  useEffect(() => {
    let ws: WebSocket | null = null;
    setUnReadCount(unread_count);
    async function fetchData() {
      const url = joinUrlParts(websocket_url, `ws?username=${username}`);
      ws = new WebSocket(url); // client side
      ws.onopen = () => {
        if (ws) {
          ws.send(`${access_token}`);
        }
      };

      ws.onmessage = (event) => {
        const obj = JSON.parse(event.data);
        setUnReadCount(obj.unread_count);
        // Assuming you have a state setter for the unread count
        // setUnReadCount(obj.unread_count);
      };

      ws.onclose = (event) => {
        // Reconnect on unexpected closure (you can add more conditions to handle specific closure codes)
        if (!event.wasClean) {
          reconnectRef.current += 1; // 使用ref的current属性修改reconnect值
          if (reconnectRef.current < 5) {
            logger.warn(
              `websocket connection closed unexpectedly, attempting to reconnect...:${reconnectRef.current}`
            );
            fetchData();
          } else {
            logger.error(
              "websocket connection closed unexpectedly, already try to reconnect 5 times"
            );
            toast(t("websocket_closed"), {
              type: "warning",
              autoClose: 2500,
            });
          }
        }
      };
    }

    fetchData();

    return () => {
      if (ws) {
        ws.close();
      }
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [access_token]);

  return (
    <div
      className="flex items-center"
      onClick={() => {
        resetUnreadCount();
        setMutationReadNotification(!mutationReadNotification);
        setNotiDialogOpen(!notiDialogOpen);
        setUnReadCount(0);
      }}
    >
      <Badge count={unReadCount} className="text-slate-700 dark:text-slate-200">
        <MdOutlineNotificationsNone className="block w-6 h-6  hover:text-sky-600 dark:hover:text-sky-400 cursor-pointer" />
      </Badge>
    </div>
  );
}
