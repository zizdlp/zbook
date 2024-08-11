import { fetchServerWithAuthWrapper } from "@/fetchs/server_with_auth";
import Pagination from "@/components/tables/pagination";
import IsEmpty from "@/components/IsEmpty";
import { FetchServerWithAuthWrapperEndPoint } from "@/fetchs/server_with_auth_util";
import { headers } from "next/headers";
import ListElementContainer from "./ListElementContainer";
import { ListDataType } from "@/fetchs/model";
import GlobalDeleteCommentDialog from "../comments/GlobalDeleteCommentDialog";
import NotFoundDemo from "../NotFoundDemo";
import { FetchError } from "@/fetchs/util";
export default async function ListElementWrapper({
  authname,
  username,
  query,
  currentPage,
  repo_name,
  listType,
}: {
  authname: string;
  username: string;
  query: string;
  currentPage: number;
  repo_name: string;
  listType: ListDataType;
}) {
  try {
    const xforward = headers().get("x-forwarded-for") ?? "";
    const agent = headers().get("User-Agent") ?? "";
    let req_counts = {
      username: username,
      query: query,
    };
    let req_elements = {
      username: username,
      page_id: currentPage,
      page_size: 10,
      query: query,
    };
    let data_counts;
    let data_elements;
    // 根据不同的 listType 设置不同的 endpoint
    if (listType === ListDataType.LIST_USER_FOLLOWER) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_FOLLOER_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_FOLLOWER,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_USER_FOLLOWING) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_FOLLOWING_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_FOLLOWING,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_PUBLIC_REPO) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_REPO_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_REPO,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_USER_REPO) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint:
          FetchServerWithAuthWrapperEndPoint.GET_LIST_USER_OWN_REPO_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_USER_OWN_REPO,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_USER_FAVORITE) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint:
          FetchServerWithAuthWrapperEndPoint.GET_LIST_USER_LIKE_REPO_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_USER_LIKE_REPO,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_ADMIN_USER) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_USER_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_USER,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_ADMIN_SESSION) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_SESSION_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_SESSION,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_REPO_VISI) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_REPO_VISIBILITY_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: {
          username: username,
          repo_name: decodeURIComponent(repo_name),
        },
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_REPO_VISIBILITY,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: {
          username: username,
          repo_name: decodeURIComponent(repo_name),
          query: query,
          page_id: currentPage,
          page_size: 10,
        },
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_ADMIN_REPO) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_REPO_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_REPO,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_ADMIN_COMMENT) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else if (listType === ListDataType.LIST_ADMIN_COMMENT_REPORT) {
      data_counts = await fetchServerWithAuthWrapper({
        endpoint:
          FetchServerWithAuthWrapperEndPoint.GET_LIST_COMMENT_REPORT_COUNT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_counts,
      });
      if (data_counts.error) {
        throw new FetchError(data_counts.message, data_counts.status);
      }
      data_elements = await fetchServerWithAuthWrapper({
        endpoint: FetchServerWithAuthWrapperEndPoint.LIST_COMMENT_REPORT,
        xforward: xforward,
        agent: agent,
        tags: [],
        values: req_elements,
      });
      if (data_elements.error) {
        throw new FetchError(data_elements.message, data_elements.status);
      }
    } else {
      throw new Error("Invalid listType");
    }
    let elements;
    const totalPages = Math.ceil(data_counts.count / 10) || 1;
    elements = data_elements.elements || [];
    return (
      <IsEmpty is_empty={elements.length == 0} listType={listType}>
        <GlobalDeleteCommentDialog />
        <div className="grid lg:grid-cols-2 gap-4 grid-cols-1">
          {elements.map((model: any, index: number) => (
            <ListElementContainer
              key={index}
              model={model}
              listType={listType}
              authname={authname}
            />
          ))}
        </div>
        {query === "" && totalPages > 1 && (
          <div className="mt-5 flex w-full justify-center">
            <Pagination totalPages={totalPages} pageKey={"page"} />
          </div>
        )}
      </IsEmpty>
    );
  } catch (error) {
    return <NotFoundDemo />;
  }
}
