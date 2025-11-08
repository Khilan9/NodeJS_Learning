
@main
def hello() = {
  class Tutorial_02_Equality_Test extends FlatSpec with Matchers {
    behavior of "DonutStore class"

    "favourite donut" should "match vanilla donut" in {
      val donutStore = new DonutStore()
      donutStore.favouriteDonut() shouldEqual "vanilla donut"
      donutStore.favouriteDonut() === "vanilla donut"
      donutStore.favouriteDonut() should not equal "plain donut"
      donutStore.favouriteDonut() should not be "plain donut"
      donutStore.favouriteDonut() !== "Plain donut"
    }
  }

  // The DonutStore class which we are testing using ScalaTest
  class DonutStore {
    def favouriteDonut(): String = "vanilla donut"
  }

  //  import scala.concurrent.ExecutionContext.Implicits.global
//  import scala.concurrent.Future
//  import scala.util.{Failure, Success}
//
//  def donutStock(donut: String): Future[Int] = Future {
//    if (donut == "vanilla donut") 10
//    else throw new IllegalStateException("Out of stock")
//  }
//
//  donutStock("vanilla donuts")
//    .recover { case e: IllegalStateException if e.getMessage == "Out of stock" => 0 }
//    .onComplete {
//      case Success(donutStock) => println(s"Results $donutStock")
//      case Failure(e) => println(s"Error processing future operations, error = ${e.getMessage}")
//    }
//
//  Thread.sleep(3000)
}
