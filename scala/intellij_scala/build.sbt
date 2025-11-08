ThisBuild / version := "0.1.0-SNAPSHOT"

ThisBuild / scalaVersion := "3.3.5"

lazy val root = (project in file("."))
  .settings(
    name := "intellij_scala"
  )

libraryDependencies ++= Seq("org.scalatest" %% "scalatest" % "3.0.1" % Test)
