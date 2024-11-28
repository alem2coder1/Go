"use client";
import Card from "@/components/ui/card-snippet";
import ValidatingMultipleRules from "./validating-multiple-rules";

const InputPage = () => {
  return (
    <div className="space-y-5">
      <Card title="Add User">
        <ValidatingMultipleRules />
      </Card>
    </div>
  );
};

export default InputPage;
