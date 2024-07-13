import "katex/dist/katex.min.css";
import Latex from "react-latex-next";
interface MathInlineProps {
  children: React.ReactNode;
}

const MathInline: React.FC<MathInlineProps> = ({ children }) => (
  <Latex>{children?.toLocaleString() ?? ""}</Latex>
);

export default MathInline;
