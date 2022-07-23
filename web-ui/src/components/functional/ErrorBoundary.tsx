import { Component, ErrorInfo, FC, ReactNode } from "react";
import { HttpError } from "src/utils/http";

type Props = {
  children: ReactNode;
  FallbackComponent: () => JSX.Element;
};

type State = {
  hasError: boolean;
};

export class ErrorBoundary extends Component<Props, State> {
  public state: State = {
    hasError: false,
  };

  public static getDerivedStateFromError(_: Error): State {
    // Update state so the next render will show the fallback UI.
    return { hasError: true };
  }

  public componentDidCatch(error: Error, errorInfo: ErrorInfo) {
    // TODO: sentry capture
    // console.error("Uncaught error:", error, errorInfo);
  }

  public render() {
    if (this.state.hasError) {
      return this.props.FallbackComponent();
    }

    return this.props.children;
  }
}
