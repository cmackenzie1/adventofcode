package twentytwenty

object ResourceLoader {
    fun loadText(name: String): String = ResourceLoader::class.java.getResource(name).readText()
}