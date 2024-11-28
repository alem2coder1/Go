"use client";

import Card from "@/components/ui/card-snippet";
import SimpleTable from "./simple-table";
import { Button } from "@/components/ui/button";
import { useNavigate } from 'react-router-dom';
import {getRoleFromJwt} from "./jwtUtils"
import CheckboxWithAction from "./checkbox-with-action";
import SelectionOperation from "./selection-operation";
import ColumnSticky from "./column-sticky";
const TailwindUiTable = () => {
    const role = getRoleFromJwt();
  return (
    <div className=" space-y-6">

        <Card title="пайдаланушылар тізімі">
            <div className="flex flex-wrap items-center gap-4 mb-1">
                <div className="flex-1">
                    <h3 className="text-xl font-medium text-default-700 mb-2">Админ</h3>
                </div>
                <div className="flex-none">
                    {role === 'superadmin' && (
                        <a type="button" href="/add-user">add-user</a>
                    )}
                </div>
            </div>
            <CheckboxWithAction/>
        </Card>
    </div>
  );
};

export default TailwindUiTable;
