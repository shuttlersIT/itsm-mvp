import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

type inputLabelPropsType = {
  label: string;
  className: string;
};

export function InputWithLabel(props: inputLabelPropsType) {
  const { label, className } = props;
  return (
    <div className="grid w-full max-w-sm items-center gap-1.5">
      <Label htmlFor="email">{label}</Label>
      <Input
        type="email"
        id="email"
        placeholder={label}
        className={className}
      />
    </div>
  );
}
