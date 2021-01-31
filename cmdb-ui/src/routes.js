/*!

=========================================================
* Material Dashboard React - v1.9.0
=========================================================

* Product Page: https://www.creative-tim.com/product/material-dashboard-react
* Copyright 2020 Creative Tim (https://www.creative-tim.com)
* Licensed under MIT (https://github.com/creativetimofficial/material-dashboard-react/blob/master/LICENSE.md)

* Coded by Creative Tim

=========================================================

* The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

*/
// @material-ui/icons
import Dashboard from "@material-ui/icons/Dashboard";
import Person from "@material-ui/icons/Person";
import LibraryBooks from "@material-ui/icons/LibraryBooks";
import BubbleChart from "@material-ui/icons/BubbleChart";
import LocationOn from "@material-ui/icons/LocationOn";
import Notifications from "@material-ui/icons/Notifications";
import Unarchive from "@material-ui/icons/Unarchive";
// core components/views for Admin layout
import DashboardPage from "./views/Dashboard/Dashboard.js";
import ApplicationList from "./views/ApplicationList/ApplicationList";
import UserProfile from "./views/UserProfile/UserProfile.js";
import TableList from "./views/TableList/TableList.js";
import Typography from "./views/Typography/Typography.js";
import Icons from "./views/Icons/Icons.js";
import Maps from "./views/Maps/Maps.js";
import NotificationsPage from "./views/Notifications/Notifications.js";
import CmdbDoc from "./views/CmdbDoc/CmdbDoc.js";
import EnvironmentList from "./views/EnvironmentList/EnvironmentList";

const dashboardRoutes = [
  {
    path: "/dashboard",
    name: "Dashboard",
    rtlName: "لوحة القيادة",
    icon: Dashboard,
    component: DashboardPage,
    layout: "/admin"
  },
  {
    path: "/applications",
    name: "Applications",
    rtlName: "قائمة الجدول",
    icon: "content_paste",
    component: ApplicationList,
    layout: "/admin"
  },
  {
    path: "/environments",
    name: "Environments",
    rtlName: "قائمة الجدول",
    icon: "content_paste",
    component: EnvironmentList,
    layout: "/admin"
  },
  {
    path: "/user",
    name: "User Profile",
    rtlName: "ملف تعريفي للمستخدم",
    icon: Person,
    component: UserProfile,
    layout: "/admin"
  },
  {
    path: "/table",
    name: "Table List",
    rtlName: "قائمة الجدول",
    icon: "content_paste",
    component: TableList,
    layout: "/admin"
  },
  {
    path: "/typography",
    name: "Typography",
    rtlName: "طباعة",
    icon: LibraryBooks,
    component: Typography,
    layout: "/admin"
  },
  {
    path: "/icons",
    name: "Icons",
    rtlName: "الرموز",
    icon: BubbleChart,
    component: Icons,
    layout: "/admin"
  },
  {
    path: "/maps",
    name: "Maps",
    rtlName: "خرائط",
    icon: LocationOn,
    component: Maps,
    layout: "/admin"
  },
  {
    path: "/notifications",
    name: "Notifications",
    rtlName: "إخطارات",
    icon: Notifications,
    component: NotificationsPage,
    layout: "/admin"
  },
  {
    path: "/cmdb-doc",
    name: "CMDB Doc",
    rtlName: "التطور للاحترافية",
    icon: Unarchive,
    component: CmdbDoc,
    layout: "/admin"
  }
];

export default dashboardRoutes;
