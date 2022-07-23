import { FC, ReactNode } from "react";
import { SWRConfig as _SWRConfig } from "swr";
import { HttpError } from "src/utils/http";

type Props = {
  children: ReactNode;
};

export const SWRConfig: FC<Props> = (props) => {
  const { children } = props;
  return (
    <_SWRConfig
      value={{
        suspense: true,
        fetcher: async (uri, init) => {
          const res = await fetch(
            `https://jsonplaceholder.typicode.com${uri}`,
            init
          );
          if (res.ok) {
            return await res.json();
          }
          const info = await res.json();
          const error = new HttpError(
            "An error occurred while fetching the data.",
            info,
            res.status
          );
          throw error;
        },
      }}
    >
      {children}
    </_SWRConfig>
  );
};
