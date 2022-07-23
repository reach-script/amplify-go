import { Component, ErrorInfo, FC, ReactNode } from "react";
import { HttpError } from "src/utils/http";

type Props = {
  children: ReactNode;
  FallbackComponent: () => JSX.Element;
};

type State = {
  hasError: boolean;
  error: HttpError | null;
};

export class HttpErrorBoundary extends Component<Props, State> {
  public state: State = {
    hasError: false,
    error: null,
  };

  public static getDerivedStateFromError(error: HttpError): State {
    return { hasError: true, error };
  }

  public componentDidCatch(error: HttpError, errorInfo: ErrorInfo) {
    if ("status" in error) {
      // TODO: sentry capture
      // console.error("Uncaught error:", error, errorInfo);
      return;
    }
    // NOTE: http errorでない場合は例外をスロー
    throw error;
  }

  public render() {
    if (!this.state.hasError) {
      return this.props.children;
    }
    if (!this.state.error) return this.props.FallbackComponent();
    switch (this.state.error.status) {
      case 404: {
        return <h3>データが見つかりませんでした</h3>;
      }
      case 500: {
        return <h3>エラーが発生しました</h3>;
      }
      default: {
        return this.props.FallbackComponent();
      }
    }
  }
}
