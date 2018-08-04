// import DashboardLayout from '../views/DashboardLayout.vue'
// GeneralViews
import NotFound from '../components/GeneralViews/NotFoundPage.vue'
// Dashboard pages
import Overview from 'src/components/Dashboard/Views/Dashboard/Overview.vue'
import Stats from 'src/components/Dashboard/Views/Dashboard/Stats.vue'

// Pages
import User from 'src/components/Dashboard/Views/Pages/UserProfile.vue'
import TimeLine from 'src/components/Dashboard/Views/Pages/TimeLinePage.vue'
// import Login from 'src/components/Dashboard/Views/Pages/Login.vue'

import Register from 'src/components/Dashboard/Views/Pages/Register.vue'
import Lock from 'src/components/Dashboard/Views/Pages/Lock.vue'

// Components pages
import Buttons from 'src/components/Dashboard/Views/Components/Buttons.vue'
import GridSystem from 'src/components/Dashboard/Views/Components/GridSystem.vue'
import Panels from 'src/components/Dashboard/Views/Components/Panels.vue'
import SweetAlert from 'src/components/Dashboard/Views/Components/SweetAlert.vue'
import Notifications from 'src/components/Dashboard/Views/Components/Notifications.vue'
import Icons from 'src/components/Dashboard/Views/Components/Icons.vue'
import Typography from 'src/components/Dashboard/Views/Components/Typography.vue'

// Forms pages
import RegularForms from 'src/components/Dashboard/Views/Forms/RegularForms.vue'
import ExtendedForms from 'src/components/Dashboard/Views/Forms/ExtendedForms.vue'
import ValidationForms from 'src/components/Dashboard/Views/Forms/ValidationForms.vue'
import Wizard from 'src/components/Dashboard/Views/Forms/Wizard.vue'

// TableList pages
import RegularTables from 'src/components/Dashboard/Views/Tables/RegularTables.vue'
import ExtendedTables from 'src/components/Dashboard/Views/Tables/ExtendedTables.vue'
// import PaginatedTables from 'src/components/Dashboard/Views/Tables/PaginatedTables.vue'
// Maps pages
import GoogleMaps from 'src/components/Dashboard/Views/Maps/GoogleMaps.vue'
import FullScreenMap from 'src/components/Dashboard/Views/Maps/FullScreenMap.vue'
import VectorMaps from 'src/components/Dashboard/Views/Maps/VectorMapsPage.vue'

// Calendar
// import Calendar from 'src/components/Dashboard/Views/Calendar/CalendarRoute.vue'
// Charts
import Charts from 'src/components/Dashboard/Views/Charts.vue'

let componentsMenu = {
  path: '/components',
  component: DashboardLayout,
  redirect: '/components/buttons',
  children: [
    {
      path: 'buttons',
      name: 'Buttons',
      component: Buttons
    },
    {
      path: 'grid-system',
      name: 'Grid System',
      component: GridSystem
    },
    {
      path: 'panels',
      name: 'Panels',
      component: Panels
    },
    {
      path: 'sweet-alert',
      name: 'Sweet Alert',
      component: SweetAlert
    },
    {
      path: 'notifications',
      name: 'Notifications',
      component: Notifications
    },
    {
      path: 'icons',
      name: 'Icons',
      component: Icons
    },
    {
      path: 'typography',
      name: 'Typography',
      component: Typography
    }

  ]
}
let formsMenu = {
  path: '/forms',
  component: DashboardLayout,
  redirect: '/forms/regular',
  children: [
    {
      path: 'regular',
      name: 'Regular Forms',
      component: RegularForms
    },
    {
      path: 'extended',
      name: 'Extended Forms',
      component: ExtendedForms
    },
    {
      path: 'validation',
      name: 'Validation Forms',
      component: ValidationForms
    },
    {
      path: 'wizard',
      name: 'Wizard',
      component: Wizard
    }
  ]
}

let tablesMenu = {
  path: '/table-list',
  component: DashboardLayout,
  redirect: '/table-list/regular',
  children: [
    {
      path: 'regular',
      name: 'Regular Tables',
      component: RegularTables
    },
    {
      path: 'extended',
      name: 'Extended Tables',
      component: ExtendedTables
    },
    {
      path: 'paginated',
      name: 'Paginated Tables',
      component: PaginatedTables
    }]
}

let mapsMenu = {
  path: '/maps',
  component: DashboardLayout,
  redirect: '/maps/google',
  children: [
    {
      path: 'google',
      name: 'Google Maps',
      component: GoogleMaps
    },
    {
      path: 'full-screen',
      name: 'Full Screen Map',
      component: FullScreenMap
    },
    {
      path: 'vector-map',
      name: 'Vector Map',
      component: VectorMaps
    }
  ]
}

let pagesMenu = {
  path: '/pages',
  component: DashboardLayout,
  redirect: '/pages/user',
  children: [
    {
      path: 'user',
      name: 'User Page',
      component: User
    },
    {
      path: 'timeline',
      name: 'Timeline Page',
      component: TimeLine
    }
  ]
}

let loginPage = {
  path: '/login',
  name: 'Login',
  component: Login
}

let registerPage = {
  path: '/register',
  name: 'Register',
  component: Register
}

let lockPage = {
  path: '/lock',
  name: 'Lock',
  component: Lock
}
import DashboardLayout from 'src/app/DashboardLayout.vue'
import Login from 'src/app/views/Login.vue'
import PaginatedTables from 'src/app/components/PaginatedTables.vue'
import Branches from '@/app/views/branches/Branches';
import BranchForm from '@/app/views/branches/Form';
import Modules from '@/app/views/modules/Modules';
import ModuleForm from '@/app/views/modules/Form';
import Classes from '@/app/views/classes/Classes';
import ClassForm from '@/app/views/classes/Form';
import Tutors from '@/app/views/users/Tutors';
import TutorForm from '@/app/views/users/TutorForm';
import Students from '@/app/views/users/Students';
import StudentForm from '@/app/views/users/StudentForm';
import Calendar from '@/app/views/calendar/CalendarRoute.vue'

const routes = [
  {
    path: '/sign',
    component: Login,
  },
  {
    path: '/',
    component: DashboardLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: '',
        name: 'Dashboard Page',
        component: Overview
      },
    ]
  },
  {
    path: '/dashboard',
    component: DashboardLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: '',
        name: 'Dashboard Page',
        component: Overview
      },
    ]
  },
  {
    path: '/activites',
    component: DashboardLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: '',
        name: 'Dashboard Page',
        component: TimeLine
      },
    ]
  },
  {
    path: '/schedules',
    component: DashboardLayout,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: '',
        component: Classes,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'create',
        component: ClassForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: ':id/edit',
        component: ClassForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'calendar',
        component: Calendar
      },
    ]
  },
  {
    path: '/admin',
    component: DashboardLayout,
    redirect: '/admin/modules',
    children: [
      {
        path: 'branches',
        component: Branches,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'branches/create',
        component: BranchForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'branches/:id/edit',
        component: BranchForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'modules',
        component: Modules,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'modules/create',
        component: ModuleForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'modules/:id/edit',
        component: ModuleForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'tutors',
        component: Tutors,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'tutors/create',
        component: TutorForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'tutors/:id/edit',
        component: TutorForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'students',
        component: Students,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'students/create',
        component: StudentForm,
        meta: {
          requiresAuth: true,
        }
      },
      {
        path: 'students/:id/edit',
        component: StudentForm,
        meta: {
          requiresAuth: true,
        }
      },
    ],
    meta: {
      requiresAuth: true,
    }
  },
  // {
  //   path: '/',
  //   component: DashboardLayout,
  //   redirect: '/admin/overview',
  //   children: [
  //     {
  //       path: 'calendar',
  //       name: 'Calendar',
  //       component: Calendar
  //     },
  //     {
  //       path: 'charts',
  //       name: 'Charts',
  //       component: Charts
  //     }
  //   ]
  // },
  // componentsMenu,
  // formsMenu,
  // tablesMenu,
  // mapsMenu,
  // pagesMenu,
  // loginPage,
  // registerPage,
  // lockPage,
  // {
  //   path: '/admin',
  //   component: DashboardLayout,
  //   redirect: '/admin/overview',
  //   children: [
  //     {
  //       path: 'overview',
  //       name: 'Overview',
  //       component: Overview
  //     },
  //     {
  //       path: 'stats',
  //       name: 'Stats',
  //       component: Stats
  //     }
  //   ]
  // },
  { path: '*', component: NotFound }
]

/**
 * Asynchronously load view (Webpack Lazy loading compatible)
 * The specified component must be inside the Views folder
 * @param  {string} name  the filename (basename) of the view to load.
 function view(name) {
   var res= require('../components/Dashboard/Views/' + name + '.vue');
   return res;
};**/

export default routes
