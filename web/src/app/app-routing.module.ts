import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TravelTabsComponent } from "./travel-tabs/travel-tabs.component";


const routes: Routes = [
  {
    path: 'travels', component: TravelTabsComponent,
  },
  {
    path: '**', pathMatch: "full", redirectTo: 'travels',
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
