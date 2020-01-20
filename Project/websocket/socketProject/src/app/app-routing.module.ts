import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { AppComponent } from './app.component';
import { SocketService } from "./socket.service";


const routes: Routes = [];

@NgModule({
  imports: [RouterModule.forRoot(routes),BrowserModule,FormsModule, HttpModule],
  exports: [RouterModule],
  declarations: [AppComponent],
  providers: [SocketService],
  bootstrap: [AppComponent]
})
export class AppRoutingModule { }
