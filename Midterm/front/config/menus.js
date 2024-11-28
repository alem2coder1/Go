import {
  Application,
  Chart,
  Components,
  DashBoard,
  Stacks2,
  Map,
  Grid,
  Files,
  Graph,
  ClipBoard,
  Cart,
  Envelope,
  Messages,
  Monitor,
  ListFill,
  Calendar,
  Flag,
  Book,
  Note,
  ClipBoard2,
  Note2,
  Note3,
  BarLeft,
  BarTop,
  ChartBar,
  PretentionChartLine,
  PretentionChartLine2,
  Google,
  Pointer,
  Map2,
  MenuBar,
  Icons,
  ChartArea,
  Building,
  Building2,
  Sheild,
  Error,
  Diamond,
  Heroicon,
  LucideIcon,
  CustomIcon,
  Mail,
} from "@/components/svg";
import {UserSign} from "../components/svg";

export const menusConfig = {
  mainNav: [

    {
      title: "Tables",
      icon: Grid,
      child: [
        {
          title: "Simple Table",
          href: "/simple-table",
          icon: BarLeft,
        },
        {
          title: "tailwindui table",
          href: "/tailwindui-table",
          icon: BarLeft,
        },
        {
          title: "Data Table",
          href: "/data-table",
          icon: BarTop,
        },
      ],
    },

  ],
  sidebarNav: {
    modern: [
      {
        title: "User",
        icon: Grid,
        child: [
          {
            title: "user-list",
            href: "/tailwindui-table",
            icon: BarLeft,
          },
          {
            title: "edit-user",
            href: "/edit-user",
            icon: UserSign,
          },
          {
            title: "add-user",
            href: "/add-user",
            icon: UserSign,
          },
        ],
      },

    ],
    classic: [
      {
        isHeader: true,
        title: "menu",
      },
      {
        title: "User",
        icon: Grid,
        child: [
          {
            title: "user-list",
            href: "/tailwindui-table",
            icon: BarLeft,
          },
          {
            title: "edit-user",
            href: "/edit-user",
            icon: UserSign,
          },
          {
            title: "add-user",
            href: "/add-user",
            icon: UserSign,
          },
        ],
      },

    ],
  },
};
