using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Pera.Web.Models;

using Pera.Extension;

namespace Pera.Web.Controllers
{
    public class HomeController : Controller
    {
        public IActionResult Index()
        {
            Person person=new Person();
            person.FirstName="David";
            person.LastName="Silva";
            return View(person);
        }

        public IActionResult About()
        {
            ViewData["Message"] = "Your application description page.";

            return View();
        }

        public IActionResult Contact()
        {
            ViewData["Message"] = "Your contact page.";

            return View();
        }

        public IActionResult Error()
        {
            return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
        }
    }
}
