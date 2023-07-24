package org.example;

import org.farng.mp3.MP3File;
import org.farng.mp3.id3.AbstractID3v2;

import java.io.File;

// Press Shift twice to open the Search Everywhere dialog and type `show whitespaces`,
// then press Enter. You can now see whitespace characters in your code.
public class Main {
    public static void main(String[] args) {
        File srcFile = new File("/home/tonni/Music/audio.mp3");
        println("Reading the file");
        try{
            MP3File mp3File = new MP3File(srcFile);
            if (mp3File.hasID3v2Tag()) {
                AbstractID3v2 tag = mp3File.getID3v2Tag();
                String title = tag.getSongTitle();
                println("Title: " + title);
                String artist = tag.getLeadArtist();
                println("Artist: " + artist);
            }  
        }
        catch (Exception e){
            System.out.println(e.getMessage());
        }
        
    }

    public static void println(String what){
        System.out.println(what);
    }
}